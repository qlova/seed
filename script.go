package seed

import (
	"fmt"
	"reflect"
	"strings"
	"net/http"
)

import "github.com/qlova/seed/script"
import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

//Set the text content of the seed.
func (seed Seed) SyncText(text *string) {
	var wrapper = func() string {
		return *text
	}
	
	seed.OnReady(func(q Script) {
		q.Javascript(`setInterval(function() {`)
			q.Get(seed).SetText(q.Call(wrapper).(qlova.String))
			for i := 0; i < q.promises; i++ {
				q.Raw("Javascript", "}; request.send();")
			}
			q.promises = 0
		q.Javascript(`}, 100)`)
	})
}


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

func (q Script) call(f interface{}, args ...qlova.Type) qlova.Type {
	if name, ok := f.(string); ok && len(args) == 0 {
		q.Raw("Javascript", language.Statement(name+`();`))
		return nil
	}
	
	var name = fmt.Sprint(f)
	
	var value = reflect.ValueOf(f)
	
	if value.Kind() != reflect.Func || value.Type().NumOut() > 1 {
		panic("Script.Call: Must pass a Go function without zero or one return values")
	}
	exports[name] = value
	
	var CallingString = `/call/`+name
	
	for i := 0; i < value.Type().NumIn(); i++ {
		switch value.Type().In(i).Kind() {
			case reflect.String:
				
				CallingString += `/_"+encodeURIComponent(`+args[i].(qlova.String).Raw()+`)+"`
				
			default:
				panic("Unimplemented: script.Run("+value.Type().String()+")")
		}
	}

	q.promises++
	q.Raw("Javascript", language.Statement(`let request = new XMLHttpRequest(); request.open("POST", "`+CallingString+`"); request.onload = function() {`))
	
	if value.Type().NumOut() == 1 {
		switch value.Type().Out(0).Kind() {
			
			case reflect.String:
				return q.Wrap(Javascript.String("this.responseText"))
			
			default:
				panic(value.Type().String()+" Unimplemented")
		}
	}
	
	return nil
}

func (q Script) Run(f interface{}, args ...qlova.Type) {
	q.call(f, args...)
}

//Export a Go function to Javascript. Don't use this for non-local apps! TODO enforce this
func (q Script) Call(f interface{}, args ...qlova.Type) qlova.Type {	
	return q.call(f, args...)
}

func callHandler(w http.ResponseWriter, r *http.Request, call string) {
	var args = strings.Split(call, "/")
	if len(args) == 0 {
		return
	}
	
	f, ok := exports[args[0]]
	if !ok {
		return
	}
	
	if len(args)-1 != f.Type().NumIn() {
		println("argument length mismatch")
		return
	}
	
	var in []reflect.Value
	for i := 0; i < f.Type().NumIn(); i++ {
		switch f.Type().In(i).Kind() {
			case reflect.String:
				
				in = append(in, reflect.ValueOf(args[i+1][1:]))
				
			default:
				println("unimplemented callHandler for "+f.Type().String())
				return
		}
	}
	
	var results = f.Call(in)
	if len(results) == 0 {
		fmt.Fprint(w, "done")
		return
	}
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
