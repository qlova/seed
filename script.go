package seed

import (
	"fmt"
	"reflect"
	"strings"
	"net/http"

	//Global ids.
	"encoding/base64"
	"math/big"
)

import "github.com/qlova/seed/script"
import "github.com/qlova/seed/style/css"
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

func (q Script) newSeed(tag string) script.Seed {
	var variable = script.Unique()
	q.Javascript(`let `+variable+` = document.createElement("`+tag+`");`)
	var seed = script.Seed{
		Native: variable,
		Qlovascript: q.Script,
	}
	seed.Style = css.Style{Stylable: seed}
	return seed
}

func (q Script) New(inherit func() Seed) script.Seed {
	var parent = inherit()
	var seed = q.newSeed(parent.tag)
	return seed
} 

func (q Script) NewSeed() script.Seed {
	return q.newSeed("div")
}

func (q Script) Contains(text, match qlova.String) qlova.Bool {
	return q.Script.BoolFromLanguageType(Javascript.Bit{Expression:language.Statement(raw(text)+`.includes(`+raw(match)+`)`)})
}

/*func (q Script) After(promise script.Promise, f func(q Script)) {
	q.Javascript(promise.Raw()+".then(function() {")
	f(q)
	q.Javascript("})")
}*/

func (q Script) After(time float64, f func()) {
	q.Javascript("setTimeout(function() {")
	f()
	q.Javascript("}, "+fmt.Sprint(time)+");")
}

func (q Script) Get(seed Interface) *script.Seed {
	return &script.Seed{
		ID: seed.GetSeed().id,
		Qlovascript: q.Script,
	}
}

func (q Script) LastPage() *script.Seed {
	return &script.Seed{
		ID: `"+last_page+"`,
		Qlovascript: q.Script,
	}
}

type global string

//All globals have a unique id.
var global_id int64 = 1;

func Global() global {
	//global identification is compressed to base64 and prefixed with g_.
	var result = "g_"+base64.RawURLEncoding.EncodeToString(big.NewInt(global_id).Bytes())

	global_id++

	return global(result)
}

func (q Script) Global(name global) qlova.String {
	return q.wrap(`window.localStorage.getItem("`+string(name)+`")`)
}

func (q Script) SetGlobal(name global, value qlova.String) {
	q.Javascript(`window.localStorage.setItem("`+string(name)+`", `+raw(value)+`);`)
}

type cookie string

//All globals have a unique id.
var cookie_id int64 = 1;

func Cookie() cookie {
	//global identification is compressed to base64 and prefixed with g_.
	var result = "c_"+base64.RawURLEncoding.EncodeToString(big.NewInt(cookie_id).Bytes())

	cookie_id++

	return cookie(result)
}

func (q Script) Cookie(name cookie) qlova.String {
	return q.wrap(`getCookie("`+string(name)+`");`)
}

func (q Script) SetCookie(name cookie, value qlova.String) {
	q.Javascript(`setCookie("`+string(name)+`", `+raw(value)+`, 365);`)
}

func ToJavascript(f func(q Script)) string {
	return string(toJavascript(f))
}

func toJavascript(f func(q Script)) []byte {
	var program = qlova.Program(func(q qlova.Script) {
		var s = Script{seedScript: &seedScript{ Script:q }}
		f(s)
		for i := 0; i < s.promises; i++ {
			q.Raw("Javascript", "}; request.send();")
		}
		s.promises = 0
	})
	source := program.SourceCode(Javascript.Implementation{})
	if source.Error {
		panic(source.ErrorMessage)
	}
	
	return source.Data
}
		

func (q Script) Javascript(js string) {
	q.Raw("Javascript", language.Statement(js))
}

func (q Script) Goto(seed Seed) {
	if !seed.page {
		q.Raw("Javascript", language.Statement(`get("`+seed.id+`").enterpage();`))
		return
	}
	q.Raw("Javascript", language.Statement(`goto("`+seed.id+`");`))
}

func (q Script) SetCurrentPage(page Seed) {
	q.Javascript(`current_page = `+page.id)
}

type Element struct {
	query string
	q Script
}

func (q Script) Query(query qlova.String) Element {
	return Element{ query: raw(query), q:q }
}

func (element Element) Run(method string) {
	element.q.Raw("Javascript", language.Statement(`document.querySelector(`+element.query+`).`+method+`();`))
}

func (q Script) Alert(message script.String) {
	q.Javascript(`alert(`+raw(message)+`);`)
}

func (q Script) Back() {
	q.Javascript(`back();`)
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
	
	var StartFrom = 0;
	//The function can take an optional client as it's first argument.
	if value.Type().NumIn() > 0 && value.Type().In(0) == reflect.TypeOf(Client{}) {
		StartFrom = 1;
	}
	
	for i := StartFrom; i < value.Type().NumIn(); i++ {
		switch value.Type().In(i).Kind() {
			case reflect.String:
				
				CallingString += `/_"+encodeURIComponent(`+raw(args[i-StartFrom].(qlova.String))+`)+"`
				
			default:
				panic("Unimplemented: script.Run("+value.Type().String()+")")
		}
	}

	q.promises++
	q.Raw("Javascript", language.Statement(`let request = new XMLHttpRequest(); request.open("POST", "`+CallingString+`"); request.onload = function() {`))
	
	if value.Type().NumOut() == 1 {
		switch value.Type().Out(0).Kind() {
			
			case reflect.String:
				return q.wrap("this.responseText")
			
			default:
				panic(value.Type().String()+" Unimplemented")
		}
	}
	
	return nil
}

func (q Script) IfNextPageIs(page Seed, f func(), els ...func()) {
	if page.seed == nil {
		if len(els) > 0 {
			q.Javascript(`{`)
			els[0]()
			q.Javascript(`}`)
		}
		return
	}

	q.Javascript(`if (next_page == "`+page.id+`") {`)
	f()
	q.Javascript(`}`)
	if len(els) > 0 {
		q.Javascript(`else {`)
		els[0]()
		q.Javascript(`}`)
	}
}

func (q Script) IfLastPageWas(page Seed, f func(), els ...func()) {
	if page.seed == nil {
		if len(els) > 0 {
			q.Javascript(`{`)
			els[0]()
			q.Javascript(`}`)
		}
		return
	}

	q.Javascript(`if (last_page  == "`+page.id+`") {`)
	f()
	q.Javascript(`}`)
	if len(els) > 0 {
		q.Javascript(`else {`)
		els[0]()
		q.Javascript(`}`)
	}
}

func (q Script) Run(f interface{}, args ...qlova.Type) {
	q.call(f, args...)
}

//Export a Go function to Javascript. Don't use this for non-local apps! TODO enforce this
func (q Script) Call(f interface{}, args ...qlova.Type) qlova.Type {	
	return q.call(f, args...)
}

func callHandler(w http.ResponseWriter, r *http.Request, call string) {
	fmt.Println(r.URL)

	var args = strings.Split(call, "/")
	if len(args) == 0 {
		return
	}
	
	f, ok := exports[args[0]]
	if !ok {
		return
	}
	
	var in []reflect.Value
	
	var StartFrom = 0;
	//The function can take an optional client as it's first argument.
	if f.Type().NumIn() > 0 && f.Type().In(0) == reflect.TypeOf(Client{}) {
		StartFrom = 1;
		
		in = append(in, reflect.ValueOf(Client{client{
			Request: r,
			ResponseWriter: w, 
		}}))
		
	}
	
	if len(args)-1 != f.Type().NumIn()-StartFrom {
		println("argument length mismatch")
		return
	}
	
	
	
	
	for i := StartFrom; i < f.Type().NumIn(); i++ {
		switch f.Type().In(i).Kind() {
			case reflect.String:
				
				in = append(in, reflect.ValueOf(args[i+1-StartFrom][1:]))
				
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

type DynamicEditor script.Seed

//Open a File object.
func (editor DynamicEditor) Open() {
	
}
