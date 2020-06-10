package script

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
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/qlova/seed/js"
	"github.com/qlova/seed/user"
)

var rpcID int64 = 1

//Go calls the given go function and tries to convert arguments to Go types.
func Go(f interface{}, args ...AnyValue) Script {
	return func(q Ctx) {
		rpc(f, false, args...)(q)
	}
}

//Wait calls the given go function and waits for it to complete.
func Wait(f interface{}, args ...AnyValue) Script {
	return func(q Ctx) {
		rpc(f, true, args...)(q)
	}
}

func RPC(f interface{}, args ...AnyValue) func(q Ctx) Value {
	return rpc(f, true, args...)
}

func rpc(f interface{}, await bool, args ...AnyValue) func(q Ctx) Value {
	return func(q Ctx) Value {
		//Get a unique string reference for f.
		var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

		rpcID++

		var value = reflect.ValueOf(f)

		if value.Kind() != reflect.Func || value.Type().NumOut() > 2 {
			panic("script.Go: Must pass a Go function without zero or one return values, not a " + reflect.TypeOf(f).String())
		}
		if value.Type().NumOut() > 2 && value.Type().Out(1) != reflect.TypeOf(error(nil)) {
			panic("script.Go: Must pass a Go function with an error value as the second parameter " + reflect.TypeOf(f).String())
		}

		Exports[name] = value

		var CallingString = `/call/` + name

		var formdata = Unique()
		var variable = Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		var f = js.Function{js.NewValue(formdata + `.set`)}

		if len(args) > 0 {
			for i, arg := range args {
				switch arg.(type) {
				case AnyFile:
					q.Run(f, q.String(strconv.Itoa(i)), arg)
				default:
					q.Run(f, q.String(strconv.Itoa(i)), js.NewValue(`JSON.stringify(%v)`, arg))
				}
			}
		}

		if await {
			q([]byte(`let ` + variable + ` = await seed.request("POST", ` + formdata + `, "` + CallingString + `", false);`))
		} else {
			q([]byte(`let ` + variable + ` = seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active);`))
		}

		return js.NewValue(variable)
	}
}

var Exports = make(map[string]reflect.Value)

//Handler returns a handler for handling remote procedure calls.
func Handler(w http.ResponseWriter, r *http.Request, call string) {
	f, ok := Exports[call]
	if !ok {
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

	default:
		err := encoder.Encode(results[0].Interface())
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
		}
	}

	//This is slow for arrays.
	u.Execute(func(q Ctx) {
		q(fmt.Sprintf(`return %v;`, buffer.String()))
	})
	return
}
