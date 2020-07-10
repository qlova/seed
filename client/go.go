package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"qlova.org/seed/js"
	"qlova.org/seed/js/console"
	"qlova.org/seed/script"
	"qlova.org/seed/user"
)

//Ctx is a client ctx
type Ctx = user.Ctx

//Script instructs the client to do something.
type Script = js.Script

//If runs the provided scripts if the clients condition is true.
func If(condition Bool, do ...Script) Script {
	return js.If(condition, script.New(do...))
}

//Contextual is any type that can load itself from Ctx.
type Contextual interface {
	FromCtx(Ctx)
}

var goExports = make(map[string]reflect.Value)

//Handler returns a handler for handling remote procedure calls.
func Handler(w http.ResponseWriter, r *http.Request, call string) {
	f, ok := goExports[call]
	if !ok {
		log.Println("invalid handler ", call)
		return
	}

	var in []reflect.Value
	var u = user.CtxFromHandler(w, r)

	var StartFrom = 0
	//The function can take an optional client as it's first argument.
	if f.Type().NumIn() > 0 && f.Type().In(0) == reflect.TypeOf(user.Ctx{}) {
		StartFrom = 1

		//Make the user, the first argument.
		in = append(in, reflect.ValueOf(u))

	}

	//The function can take an optional Contextual as it's first argument.
	if f.Type().NumIn() > 0 && f.Type().In(0) == reflect.TypeOf([0]Contextual{}).Elem() {
		StartFrom = 1

		var zero = reflect.Zero(f.Type().In(0))

		zero.Interface().(Contextual).FromCtx(u)

		//Make the user, the first argument.
		in = append(in, zero)

	}

	//Parse each argument as JSON.
	for i := StartFrom; i < f.Type().NumIn(); i++ {
		var arg = u.Arg(strconv.Itoa(i - StartFrom))

		var elem reflect.Value
		var reader *io.Reader

		switch f.Type().In(i) {

		case reflect.TypeOf(user.File{}):
			elem = reflect.ValueOf(arg.File())
		case reflect.TypeOf(reader).Elem():
			elem = reflect.ValueOf(arg.File())
		case reflect.TypeOf(time.Time{}):
			t, err := time.Parse(`"2006-01-02"`, arg.String())
			if err != nil {
				t, err = time.Parse(`"2006-01"`, arg.String())
			}
			elem = reflect.ValueOf(t)
		case reflect.TypeOf(true):
			if arg.String() == "true" {
				elem = reflect.ValueOf(true)
			} else {
				elem = reflect.ValueOf(false)
			}
		case reflect.TypeOf(url.URL{}):
			s, err := strconv.Unquote(arg.String())
			if err != nil {
				log.Println(err)
				u.Report(errors.New("invalid url"))
				return
			}
			location, err := url.Parse(s)
			if err != nil {
				log.Println(err)
				u.Report(errors.New("invalid url"))
				return
			}
			elem = reflect.ValueOf(*location)
		default:
			var shell = reflect.New(f.Type().In(i)).Interface()
			if err := json.NewDecoder(strings.NewReader(arg.String())).Decode(shell); err != nil {
				log.Println("could not decode argument: ", err)
				u.Report(errors.New("invalid request"))
				return
			}

			elem = reflect.ValueOf(shell).Elem()
		}

		var ElemType, ArgType = elem.Type(), f.Type().In(i)
		if ElemType != ArgType {
			if !(ArgType.Kind() == reflect.Interface && ElemType.Implements(ArgType)) {
				log.Println("type mismatch")
				u.Report(errors.New("invalid request"))
				return
			}
		}

		in = append(in, elem)
	}

	var results = f.Call(in)

	if len(results) == 0 {
		fmt.Fprintf(u.ResponseWriter(), "//ok")
		return
	}

	if err, ok := results[len(results)-1].Interface().(error); ok {
		u.Report(err)
		return
	}

	var result = results[0].Interface()

	var buffer bytes.Buffer

	var encoder = json.NewEncoder(&buffer)

	type JSONEncodable interface {
		JSONEncoder() func(interface{}) ([]byte, error)
	}

	switch v := result.(type) {
	case JSONEncodable:
		var encoder = v.JSONEncoder()

		b, err := encoder(result)
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
			return
		}
		buffer.Write(b)

	case json.Marshaler:
		encoded, err := v.MarshalJSON()
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
			return
		}
		buffer.Write(encoded)

	default:
		err := encoder.Encode(results[0].Interface())
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
		}
	}

	//This is slow for arrays.
	u.Execute(func(q script.Ctx) {
		q(fmt.Sprintf(`return %v;`, buffer.String()))
	})
	return
}

//Go requests the client to call the given Go function with the given client Values automatically converted to equivalent Go values and are passed to the given function.
//The function can optionally take a Ctx as the first argument, if so, then it is passed to the function and arguments are assigned to the following arguments.
func Go(fn interface{}, args ...Value) Script {
	return func(q script.Ctx) {
		//Get a unique string reference for f.
		var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

		rpcID++

		var value = reflect.ValueOf(fn)

		if value.Kind() != reflect.Func || value.Type().NumOut() > 2 {
			panic("script.Go: Must pass a Go function without zero or one return values, not a " + reflect.TypeOf(fn).String())
		}
		if value.Type().NumOut() > 2 && value.Type().Out(1) != reflect.TypeOf(error(nil)) {
			panic("script.Go: Must pass a Go function with an error value as the second parameter " + reflect.TypeOf(fn).String())
		}

		goExports[name] = value

		var CallingString = `/go/` + name

		var formdata = script.Unique()
		var variable = script.Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		var f = js.Function{js.NewValue(formdata + `.set`)}

		if len(args) > 0 {
			for i, arg := range args {
				switch arg.(type) {
				case script.AnyFile:
					q.Run(f, q.String(string('a'+rune(i))), arg)
				default:
					q.Run(f, q.String(string('a'+rune(i))), js.NewValue(`JSON.stringify(%v)`, arg))
				}
			}
		}

		//if await {
		//	q([]byte(`let ` + variable + ` = await seed.request("POST", ` + formdata + `, "` + CallingString + `", false);`))
		//} else {
		q([]byte(`let ` + variable + ` = seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active);`))
		//}

		//return js.NewValue(variable)
	}
}

var rpcID int64

func Download(fn interface{}, args ...Value) Script {
	return func(q script.Ctx) {
		//Get a unique string reference for f.
		var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

		rpcID++

		var value = reflect.ValueOf(fn)

		if value.Kind() != reflect.Func || value.Type().NumOut() > 2 {
			panic("script.Go: Must pass a Go function without zero or one return values, not a " + reflect.TypeOf(fn).String())
		}
		if value.Type().NumOut() > 2 && value.Type().Out(1) != reflect.TypeOf(error(nil)) {
			panic("script.Go: Must pass a Go function with an error value as the second parameter " + reflect.TypeOf(fn).String())
		}

		goExports[name] = value

		var CallingString = `/call/` + name + `?`

		var formdata = script.Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		var f = js.Function{js.NewValue(formdata + `.set`)}

		if len(args) > 0 {
			for i, arg := range args {
				switch arg.(type) {
				default:
					q.Run(f, q.String(string('a'+rune(i))), js.NewValue(`JSON.stringify(%v)`, arg))
				}
			}
		}

		q(console.Log(NewString(CallingString).GetString().Plus(
			js.String{js.NewValue(`new URLSearchParams(%v).toString()`, js.NewValue(formdata))},
		)))

		q(js.Func(`seed.download`).Run(NewString(""), NewString(CallingString).GetString().Plus(
			js.String{js.NewValue(`new URLSearchParams(%v).toString()`, js.NewValue(formdata))},
		)))
	}
}
