package script

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
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
		RPC(f, args...)
	}
}

func RPC(f interface{}, args ...AnyValue) func(q Ctx) Value {
	return func(q Ctx) Value {
		//Get a unique string reference for f.
		var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

		rpcID++

		var value = reflect.ValueOf(f)

		if value.Kind() != reflect.Func || value.Type().NumOut() > 1 {
			panic("Script.Call: Must pass a Go function without zero or one return values, not a " + reflect.TypeOf(f).String())
		}
		Exports[name] = value

		var CallingString = `/call/` + name

		var formdata = Unique()
		var variable = Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		if len(args) > 0 {
			for i, arg := range args {
				switch arg.(type) {
				case AnyFile:
					q.Run(formdata+`.set`, q.String(strconv.Itoa(i)), arg)
				default:
					q.Run(formdata+`.set`, q.String(strconv.Itoa(i)), js.NewValue(`JSON.stringify(%v)`, arg))
				}
			}
		}

		q([]byte(`let ` + variable + ` = await seed.request("POST", ` + formdata + `, "` + CallingString + `");`))

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
		default:
			var shell = reflect.New(f.Type().In(i)).Interface()
			if err := json.NewDecoder(strings.NewReader(arg.String())).Decode(shell); err != nil {
				log.Println(err)
				return
			}

			elem = reflect.ValueOf(shell).Elem()
		}

		var ElemType, ArgType = elem.Type(), f.Type().In(i)
		if ElemType != ArgType {
			if !(ArgType.Kind() == reflect.Interface && ElemType.Implements(ArgType)) {
				log.Println("type mismatch")
				return
			}
		}

		in = append(in, elem)
	}

	var results = f.Call(in)

	if len(results) == 0 {
		return
	}

	if len(results) == 1 {
		if err, ok := results[0].Interface().(error); ok {
			u.Report(err)
			return
		}
		var buffer bytes.Buffer
		err := json.NewEncoder(&buffer).Encode(results[0].Interface())
		if err != nil {
			fmt.Println("rpc function could not send return value: ", err)
		}
		u.Execute(fmt.Sprintf(`return %v;`, buffer.String()))
		return
	}

	panic("rpc function with more than one return value")
}
