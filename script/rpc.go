package script

import (
	//Global ids.
	"encoding/base64"
	"fmt"
	"math/big"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/qlova/seed/user"

	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"
)

//Request is the JS code required to make Go calls.
const Request = `
function request (method, formdata, url, manual) {
	if (ServiceWorker_Registration) ServiceWorker_Registration.update();

	if (url.charAt(0) == "/") url = host+url;

	if (manual) {
			var xhr = new XMLHttpRequest();
			xhr.open(method, url);
		return xhr;
	}

	return new Promise(function (resolve, reject) {
	var xhr = new XMLHttpRequest();
	xhr.open(method, url, true);
	xhr.onload = function () {
		if (this.status >= 200 && this.status < 300) {
		resolve(xhr.response);
		} else {
		reject({
			status: this.status,
			statusText: xhr.statusText,
			response: xhr.response
		});
		}
	};
	xhr.onerror = function () {
		reject({
		status: this.status,
		statusText: xhr.statusText,
		response: xhr.response
		});
	};
	xhr.send(formdata);
	});
}
`

//Promise represents a future action that can either succeed or fail.
type Promise struct {
	expression string
	q          Script
}

//Raw returns the raw JS promise.
func (promise Promise) Raw() string {
	return promise.expression
}

//Then executes the provided function when the promise succeeds.
func (promise Promise) Then(f func()) Promise {
	promise.q.Javascript(promise.expression + ".then(function(rpc_result) {")
	f()
	promise.q.Javascript("}).catch(function(){});;")
	return promise
}

//Catch executes the provided function when the promise fails.
func (promise Promise) Catch(f func()) Promise {
	promise.q.Javascript(promise.expression + ".catch(function(rpc_result) {")
	f()
	promise.q.Javascript("});")
	return promise
}

//Args is a mapping from strings to script types.
type Args map[string]qlova.Type

//Attachable is a something that can be attached to a Go call.
type Attachable interface {
	AttachTo(request string, index int) string
}

//Attach attaches Attachables and returns an AttachCall.
func (q Script) Attach(attachables ...Attachable) Attached {
	var variable = Unique()

	q.Javascript(`var ` + variable + " = new FormData();")

	for i, attachable := range attachables {
		q.Javascript(attachable.AttachTo(variable, i+1))
	}

	return Attached{variable, q, nil}
}

//Attached has attachments and these will be passed to the Go function that is called.
type Attached struct {
	formdata string
	q        Script
	args     Args
}

//Go calls a Go function f, with args. Returns a promise.
func (c Attached) Go(f interface{}, args ...qlova.Type) Promise {
	return c.q.rpc(f, c.formdata, c.args, args...)
}

//With adds arguments to the attached call.
func (c Attached) With(args Args) Attached {
	c.args = args
	return c
}

//With adds arguments to the attached call.
func (q Script) With(args Args) Attached {
	return Attached{"", q, args}
}

var rpcID int64

func (q Script) rpc(f interface{}, formdata string, nargs Args, args ...qlova.Type) Promise {

	//Get a unique string reference for f.
	var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

	rpcID++

	var value = reflect.ValueOf(f)

	if value.Kind() != reflect.Func || value.Type().NumOut() > 1 {
		panic("Script.Call: Must pass a Go function without zero or one return values")
	}
	exports[name] = value

	var CallingString = `/call/` + name

	var StartFrom = 0
	//The function can take an optional client as it's first argument.
	if value.Type().NumIn() > 0 && value.Type().In(0) == reflect.TypeOf(user.User{}) {
		StartFrom = 1
	}

	for i := StartFrom; i < value.Type().NumIn(); i++ {
		switch value.Type().In(i).Kind() {
		case reflect.String, reflect.Int:

			CallingString += `/_"+encodeURIComponent(` + raw(args[i-StartFrom].(qlova.String)) + `)+"`

		default:
			panic("Unimplemented: script.Run(" + value.Type().String() + ")")
		}
	}

	var variable = Unique()

	if nargs != nil {
		if formdata == "" {
			formdata = Unique()
			q.Javascript(formdata + ` = new FormData();`)
		}
		for key, value := range nargs {
			q.Javascript(formdata + `.set(` + strconv.Quote(key) + `, ` + value.LanguageType().Raw() + `);`)
		}
	}

	q.Require(Request)
	q.Raw("Javascript", language.Statement(`let `+variable+` = request("POST", `+formdata+`, "`+CallingString+`");`))

	return Promise{variable, q}
}

//ReturnValue can be used to access the Go return value as a string.
//Only works inside a Promise callback, otherwise behaviour is undefined.
func (q Script) ReturnValue() qlova.String {
	return q.wrap("rpc_result")
}

//Error can be used to access the Go return error as a string.
//Only works inside a Promise callback, otherwise behaviour is undefined.
func (q Script) Error() qlova.String {
	return q.wrap("rpc_result.response")
}

var exports = make(map[string]reflect.Value)

func (q Script) call(f interface{}, args ...qlova.Type) qlova.Value {
	if name, ok := f.(string); ok && len(args) == 0 {
		q.Raw("Javascript", language.Statement(name+`();`))
		return qlova.Value{}
	}

	var name = fmt.Sprint(f)

	var value = reflect.ValueOf(f)

	if value.Kind() != reflect.Func || value.Type().NumOut() > 1 {
		panic("Script.Call: Must pass a Go function without zero or one return values")
	}
	exports[name] = value

	var CallingString = `/call/` + name

	var StartFrom = 0
	//The function can take an optional client as it's first argument.
	if value.Type().NumIn() > 0 && value.Type().In(0) == reflect.TypeOf(user.User{}) {
		StartFrom = 1
	}

	for i := StartFrom; i < value.Type().NumIn(); i++ {
		switch value.Type().In(i).Kind() {
		case reflect.String:

			CallingString += `/_"+encodeURIComponent(` + raw(args[i-StartFrom].(qlova.String)) + `)+"`

		default:
			panic("Unimplemented: script.Run(" + value.Type().String() + ")")
		}
	}

	q.Require(Request)
	q.Raw("Javascript", language.Statement(`let request = new XMLHttpRequest(); request.open("POST", "`+CallingString+`"); request.onload = function() {`))

	if value.Type().NumOut() == 1 {
		switch value.Type().Out(0).Kind() {

		case reflect.String:
			return q.wrap("this.responseText").Value()

		default:
			panic(value.Type().String() + " Unimplemented")
		}
	}

	return qlova.Value{}
}

//Handler returns a handler for handling remote procedure calls.
func Handler(w http.ResponseWriter, r *http.Request, call string) {
	var args = strings.Split(call, "/")
	if len(args) == 0 {
		return
	}

	f, ok := exports[args[0]]
	if !ok {
		return
	}

	var in []reflect.Value

	var u = user.User{}.FromHandler(w, r)

	var StartFrom = 0
	//The function can take an optional client as it's first argument.
	if f.Type().NumIn() > 0 && f.Type().In(0) == reflect.TypeOf(user.User{}) {
		StartFrom = 1

		in = append(in, reflect.ValueOf(u))

	}

	if len(args)-1 != f.Type().NumIn()-StartFrom {
		println("argument length mismatch")
		return
	}

	for i := StartFrom; i < f.Type().NumIn(); i++ {
		switch f.Type().In(i).Kind() {
		case reflect.String:

			in = append(in, reflect.ValueOf(args[i+1-StartFrom][1:]))

		case reflect.Int:
			var number, _ = strconv.Atoi(args[i+1-StartFrom][1:])
			in = append(in, reflect.ValueOf(number))

		default:
			println("unimplemented callHandler for " + f.Type().String())
			return
		}
	}

	var results = f.Call(in)

	u.Close()

	if len(results) == 0 {
		return
	}

	switch results[0].Kind() {

	case reflect.String:
		if results[0].Interface().(string) == "" {
			//Error
			http.Error(w, "", 500)
			return
		}
		fmt.Fprint(w, results[0].Interface())

	default:
		fmt.Println(results[0].Type().String(), " Unimplemented")
	}
}
