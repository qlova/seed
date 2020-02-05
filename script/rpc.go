package script

import (
	"log"
	"strings"

	//Global ids.
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"reflect"
	"strconv"

	"github.com/qlova/seed/user"

	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"
)

//Request is the JS code required to make Go calls.
const Request = `
function slave(response) {
	return (new Function(response))();
}

function request (method, formdata, url, manual) {

	if (window.rpc && rpc[url]) {
		slave(rpc[url](formdata));
		return;
	}

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
				resolve(slave(xhr.response));
			} else {
				slave(xhr.response);
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

//Args is a mapping from strings to script types.
type Args map[string]qlova.Type

//Attachable is a something that can be attached to a Go call.
type Attachable interface {
	AttachTo(request string, index int) string
}

//Attach attaches Attachables and returns an AttachCall.
func (q Ctx) Attach(attachables ...Attachable) Attached {
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
	q        Ctx
	args     Args
}

//Go calls a Go function f, with args. Returns a promise.
func (c Attached) Go(f interface{}, args ...qlova.Type) Promise {
	return c.q.rpc(f, c.formdata, c.args, args...)
}

//With adds arguments to the attached call.
func (c Attached) With(args Args) Attached {
	if c.args == nil {
		c.args = args
	}
	for key, value := range args {
		c.args[key] = value
	}
	return c
}

//With adds arguments to the attached call.
func (q Ctx) With(args Args) Attached {
	return Attached{"", q, args}
}

var rpcID int64 = 1

func (q Ctx) rpc(f interface{}, formdata string, nargs Args, args ...qlova.Type) Promise {
	//Get a unique string reference for f.
	var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

	rpcID++

	var value = reflect.ValueOf(f)

	if value.Kind() != reflect.Func || value.Type().NumOut() > 1 {
		panic("Script.Call: Must pass a Go function without zero or one return values")
	}
	Exports[name] = value

	var CallingString = `/call/` + name

	var variable = Unique()

	//Get all positional arguments and add them to the formdata.
	if len(args) > 0 {
		if formdata == "" || formdata == "undefined" {
			formdata = Unique()
			q.Javascript(`let ` + formdata + ` = new FormData();`)
		}

		for i, arg := range args {
			q.Javascript(`%v.set("%v", JSON.stringify(%v));`, formdata, i, arg)
		}
	}

	//Get all named arguments and add them to the formdata.
	if nargs != nil {
		if formdata == "" || formdata == "undefined" {
			formdata = Unique()
			q.Javascript(`let ` + formdata + ` = new FormData();`)
		}
		for key, value := range nargs {
			q.Javascript(formdata + `.set(` + strconv.Quote(key) + `, JSON.stringify(` + value.LanguageType().Raw() + `));`)
		}
	}

	q.Require(Request)
	q.Raw("Javascript", language.Statement(`let `+variable+` = request("POST", `+formdata+`, "`+CallingString+`");`))

	return Promise{q.Value(variable).Native(), q}
}

//ReturnValue can be used to access the Go return value as a string.
//Only works inside a Promise callback, otherwise behaviour is undefined.
func (q Ctx) ReturnValue() qlova.String {
	return q.wrap("rpc_result")
}

//Error can be used to access the Go return error as a string.
//Only works inside a Promise callback, otherwise behaviour is undefined.
func (q Ctx) Error() qlova.String {
	return q.wrap("rpc_result.response")
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
