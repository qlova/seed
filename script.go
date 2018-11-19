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

func (q Script) Query(query script.String) Element {
	return Element{ query:dynamicString(query), q:q }
}

func (element Element) Run(method string) {
	element.q.Raw("Javascript", language.Statement(`document.querySelector(`+element.query+`).`+method+`();`))
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
func (q Script) Call(f interface{}) script.Type {
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

type DynamicSeed struct {
	id string
	q script.Script
}

//Convert a Qlovascript string into a Javascript string.
func dynamicString(s script.String) string {
	var text string
	if s.Literal == nil {
		text = string(s.String.(Javascript.String))
	} else {
		text = `"`+*s.Literal+`"`
	}
	return text
}

func (seed DynamicSeed) SetText(s script.String) {
	seed.q.Raw("Javascript", language.Statement(`get("`+seed.id+`").textContent = `+dynamicString(s)+`;`))
}

func (seed DynamicSeed) SetLeft(s script.String) {
	seed.q.Raw("Javascript", language.Statement(`get("`+seed.id+`").style.left = `+dynamicString(s)+`;`))
}

func (seed DynamicSeed) SetDisplay(s script.String) {
	seed.q.Raw("Javascript", language.Statement(`get("`+seed.id+`").style.display = `+dynamicString(s)+`;`))
}

func (seed DynamicSeed) SetVisible() {
	seed.q.Raw("Javascript", language.Statement(`get("`+seed.id+`").style.display = "block";`))
}

func (seed DynamicSeed) SetHidden() {
	seed.q.Raw("Javascript", language.Statement(`get("`+seed.id+`").style.display = "none";`))
}

func (seed DynamicSeed) Click() {
	seed.q.Raw("Javascript", language.Statement(`get("`+seed.id+`").click();`))
}

func (seed DynamicSeed) Left() script.String {
	return seed.q.Wrap(Javascript.String(`get("`+seed.id+`").style.left`)).(script.String)
}

func (seed DynamicSeed) Width() script.String {
	return seed.q.Wrap(Javascript.String(`getComputedStyle(get("`+seed.id+`")).width`)).(script.String)
}
