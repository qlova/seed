package script

import (
	"fmt"
	"reflect"
)

import "github.com/qlova/seed/user"

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"

type Promise struct {
	expression string
	q Script
}

func (promise Promise) Then(f func()) Promise {
	promise.q.Javascript(promise.expression+".then(function(rpc_result) {")
	f()
	promise.q.Javascript("})")
	return Promise{"", promise.q}
}

func (promise Promise) Catch(f func()) Promise {
	promise.q.Javascript(promise.expression+".catch(function(rpc_result) {")
	f()
	promise.q.Javascript("});")
	return promise
}

func (q Script) rpc(f interface{}, formdata string, args ...qlova.Type) Promise {

	//Get a unique string reference for f.	
	var name = fmt.Sprint(f)
	
	var value = reflect.ValueOf(f)
	
	if value.Kind() != reflect.Func || value.Type().NumOut() > 1 {
		panic("Script.Call: Must pass a Go function without zero or one return values")
	}
	exports[name] = value
	
	var CallingString = `/call/`+name
	
	var StartFrom = 0;
	//The function can take an optional client as it's first argument.
	if value.Type().NumIn() > 0 && value.Type().In(0) == reflect.TypeOf(user.User{}) {
		StartFrom = 1;
	}
	
	for i := StartFrom; i < value.Type().NumIn(); i++ {
		switch value.Type().In(i).Kind() {
			case reflect.String, reflect.Int:
				
				CallingString += `/_"+encodeURIComponent(`+raw(args[i-StartFrom].(qlova.String))+`)+"`
				
			default:
				panic("Unimplemented: script.Run("+value.Type().String()+")")
		}
	}

	var variable = Unique() 
	
	q.Raw("Javascript", language.Statement(`let `+variable+` = request("POST", `+formdata+`, "`+CallingString+`");`))
	
	return Promise{variable, q}
}

func (q Script) ReturnValue() qlova.String {
	return q.wrap("rpc_result")
}

//Call a Go function from within a script. The result is returned as a promise.
func (q Script) ServerCall(f interface{}, args ...qlova.Type) Promise {	
	return q.rpc(f, "undefined", args...)
}
