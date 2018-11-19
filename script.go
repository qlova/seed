package seed

import "fmt"
import "reflect"
import "net/http"

import "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type Script struct {
	*seedScript
}

type seedScript struct {
	script.Script
	promises int
}

func (q Script) Get(seed Seed) DynamicSeed {
	return DynamicSeed{
		id: seed.id,
		q: q.Script,
	}
}

type ExportedFunction struct {
	f reflect.Value
}

var exports = make(map[string]reflect.Value)

//Export a Go function to Javascript. Don't use this for non-local apps! TODO enforce this
func (q Script) Call(f interface{}) script.Type {
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

type DynamicSeed struct {
	id string
	q script.Script
}

func (seed DynamicSeed) SetText(s script.String) {
	var text string
	if s.Literal == nil {
		text = string(s.String.(Javascript.String))
	} else {
		text = `"`+*s.Literal+`"`
	}
	seed.q.Raw("Javascript", language.Statement(`document.getElementById("`+seed.id+`").textContent = `+text+`;`))
}

