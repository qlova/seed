package seed

import (
	"fmt"
	"reflect"
)

import qlova "github.com/qlova/script"
import "github.com/qlova/seed/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

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
	promise.q.Javascript("})")
	return promise
}

func (q Script) rpc(f interface{}, args ...qlova.Type) Promise {

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
	if value.Type().NumIn() > 0 && value.Type().In(0) == reflect.TypeOf(Client{}) {
		StartFrom = 1;
	}
	
	for i := StartFrom; i < value.Type().NumIn(); i++ {
		switch value.Type().In(i).Kind() {
			case reflect.String:
				
				CallingString += `/_"+encodeURIComponent(`+args[i-StartFrom].(qlova.ExportedString).Raw()+`)+"`
				
			default:
				panic("Unimplemented: script.Run("+value.Type().String()+")")
		}
	}

	var variable = script.Unique() 
	
	q.Raw("Javascript", language.Statement(`let `+variable+` = request("POST", "`+CallingString+`");`))
	
	return Promise{variable, q}
}

func (q Script) ReturnValue() qlova.ExportedString {
	return q.Wrap(Javascript.String("rpc_result")).(qlova.ExportedString)
}

//Call a Go function from within a script. The result is returned as a promise.
func (q Script) ServerCall(f interface{}, args ...qlova.Type) Promise {	
	return q.rpc(f, args...)
}