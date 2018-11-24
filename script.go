package seed

import (
	"fmt"
	"reflect"
	"net/http"
)

import "github.com/qlova/seed/script"
import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type Script struct {
	*seedScript
}

type seedScript struct {
	qlova.Script
	promises int
}

func (q Script) Get(seed Seed) script.Seed {
	return script.Seed{
		ID: seed.id,
		Qlovascript: q.Script,
	}
}

func toJavascript(f func(q Script)) []byte {
	var program = qlova.NewProgram(func(q qlova.Script) {
		var s = Script{seedScript: &seedScript{ Script:q }}
		f(s)
		for i := 0; i < s.promises; i++ {
			q.Raw("Javascript", "}; request.send();")
		}
		s.promises = 0
	})
	source, err := program.Source(Javascript.Language())
	if err != nil {
		panic(err)
	}
	
	return []byte(source)
}
		

func (q Script) Javascript(js string) {
	q.Raw("Javascript", language.Statement(js))
}

func (q Script) Goto(seed Seed) {
	q.Raw("Javascript", language.Statement(`goto("`+seed.id+`");`))
}

type Element struct {
	query string
	q Script
}

func (q Script) Query(query qlova.String) Element {
	return Element{ query:query.Raw(), q:q }
}

func (element Element) Run(method string) {
	element.q.Raw("Javascript", language.Statement(`document.querySelector(`+element.query+`).`+method+`();`))
}

func (q Script) Alert(message script.String) {
	q.Raw("Javascript", language.Statement(`alert(`+message.Raw()+`);`))
}

type ExportedFunction struct {
	f reflect.Value
}

var exports = make(map[string]reflect.Value)

func (q Script) Run(f interface{}) {
	if name, ok := f.(string); ok {
		q.Raw("Javascript", language.Statement(name+`();`))
		return 
	}
	
	panic("script.Run(func()): Unimplemented")
}

//Export a Go function to Javascript. Don't use this for non-local apps! TODO enforce this
func (q Script) Call(f interface{}) qlova.Type {
	if _, ok := f.(string); ok {
		panic("script.Run(string): Unimplemented")
		return nil
	}
	
	q.promises++
	
	var name = fmt.Sprint(f)
	
	var value = reflect.ValueOf(f) 
	
	if value.Kind() != reflect.Func || value.Type().NumOut() != 1 {
		panic("Script.Call: Must pass a Go function with 1 return value")
	}
	exports[name] = value
	
	
	q.Raw("Javascript", language.Statement(`let request = new XMLHttpRequest(); request.open("POST", "/call/`+name+`"); request.onload = function() {`))
	
	switch value.Type().Out(0).Kind() {
		
		case reflect.String:
			return q.Wrap(Javascript.String("this.responseText"))
		
		default:
			panic(value.Type().String()+" Unimplemented")
	}
	
	return nil
}

func callHandler(w http.ResponseWriter, r *http.Request, call string) {
	f, ok := exports[call]
	if !ok {
		return
	}
	
	var results = f.Call(nil)
	switch results[0].Kind() {
		
		case reflect.String:
			fmt.Fprint(w, results[0].Interface())
			
		default:
			fmt.Println(results[0].Type().String(), " Unimplemented")
	}
}

type DynamicEditor script.Seed

//Open a File object.
func (editor DynamicEditor) Open() {
	
}
