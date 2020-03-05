package script

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/qlova/script"
	"github.com/qlova/seed/user"
)

//Go calls the given go function and tries to convert arguments to Go types.
//To check error, call .Error() on the result.
func Go(f interface{}, args ...script.AnyValue) Script {
	return func(q Ctx) {
		q.rpc(f, args...)
	}
}

//Go calls the given go function and tries to convert arguments to Go types.
//To check error, call .Error() on the result.
func (q Ctx) Go(f interface{}, args ...script.AnyValue) Promise {
	return q.rpc(f, args...)
}

var rpcID int64 = 1

func (q Ctx) rpc(f interface{}, args ...script.AnyValue) Promise {
	//Get a unique string reference for f.
	var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

	rpcID++

	var value = reflect.ValueOf(f)

	if value.Kind() != reflect.Func || value.Type().NumOut() > 1 {
		panic("Script.Call: Must pass a Go function without zero or one return values, not a " + reflect.TypeOf(f).String())
	}
	Exports[name] = value

	var CallingString = `/call/` + name

	var variable = Unique()
	var formdata = Unique()
	q.Javascript(`let ` + formdata + ` = new FormData();`)

	//Get all positional arguments and add them to the formdata.
	if len(args) > 0 {

		for i, arg := range args {
			q.Javascript(`%v.set("%v", JSON.stringify(%v));`, formdata, i, arg.ValueFromCtx(q))
		}
	}

	q.Write([]byte(`let ` + variable + ` = seed.request("POST", ` + formdata + `, "` + CallingString + `");`))

	return q.Value(variable).Promise()
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

		var shell = reflect.New(f.Type().In(i)).Interface()
		if err := json.NewDecoder(strings.NewReader(arg.String())).Decode(shell); err != nil {
			log.Println(err)
			return
		}

		var elem = reflect.ValueOf(shell).Elem()

		if elem.Type() != f.Type().In(i) {
			log.Println("type mismatch")
			return
		}

		in = append(in, elem)
	}

	var results = f.Call(in)

	if len(results) == 0 {
		return
	}

	if len(results) == 1 {
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
