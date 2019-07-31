package script

import (
	"fmt"
	"reflect"

	"net/http"
	"strconv"
	"strings"
)

import "github.com/qlova/seed/internal"
import "github.com/qlova/seed/user"

import "github.com/qlova/seed/style/css"
import qlova "github.com/qlova/script"

import "github.com/qlova/script/language"
import Javascript "github.com/qlova/script/language/javascript"

type Script struct {
	*script
}

type script struct {
	internal.Context
	qlova.Script

	js   js
	Time time
}

func (q Script) Require(dependency string) {

	//Subdependencies.
	if dependency == Goto {
		q.Require(Get)
		q.Require(Set)
	}

	if _, ok := q.Dependencies[dependency]; ok {
		return
	}
	q.Dependencies[dependency] = struct{}{}
}

func (q Script) RawString(s qlova.String) string {
	return raw(s)
}

func raw(s qlova.String) string {
	return string(s.LanguageType().(Javascript.String).Expression)
}

func (q Script) wrap(s string) qlova.String {
	return q.StringFromLanguageType(Javascript.String{
		Expression: language.Statement(s),
	})
}

func (q Script) newSeed(tag string) Seed {
	var variable = Unique()
	q.Javascript(`let ` + variable + ` = document.createElement("` + tag + `");`)
	var seed = Seed{
		Native: variable,
		Q:      q,
	}
	seed.Style = css.Style{Stylable: seed}
	return seed
}

/*func (q Script) New(inherit func() Seed) script.Seed {
	var parent = inherit()
	var seed = q.newSeed(parent.tag)
	return seed
}*/

func (q Script) NewSeed(tag ...string) Seed {

	if len(tag) > 0 {
		return q.newSeed(tag[0])
	}

	return q.newSeed("div")
}

func (q Script) Contains(text, match qlova.String) qlova.Bool {
	return q.Script.BoolFromLanguageType(Javascript.Bit{Expression: language.Statement(raw(text) + `.includes(` + raw(match) + `)`)})
}

/*func (q Script) After(promise script.Promise, f func(q Script)) {
	q.Javascript(promise.Raw()+".then(function() {")
	f(q)
	q.Javascript("})")
}*/

func (q Script) After(time float64, f func()) {
	q.Javascript("setTimeout(function() {")
	f()
	q.Javascript("}, " + fmt.Sprint(time) + ");")
}

func (q Script) Every(time float64, f func()) {
	q.Javascript("setInterval(function() {")
	f()
	q.Javascript("}, " + fmt.Sprint(time) + ");")
}

/*func (q Script) Get(seed Interface) *script.Seed {
	return &script.Seed{
		ID: seed.GetSeed().id,
		Qlovascript: q.Script,
	}
}*/

func (q Script) LastPage() Page {
	return Page{Seed{
		ID: `"+last_page+"`,
		Q:  q,
	}}
}

func (q Script) NextPage() Page {
	return Page{Seed{
		ID: `"+next_page+"`,
		Q:  q,
	}}
}

const GetCookie = `
	function getCookie(cname) {
		var name = cname + "=";
		var decodedCookie = decodeURIComponent(document.cookie);
		var ca = decodedCookie.split(';');
		for(var i = 0; i <ca.length; i++) {
		var c = ca[i];
		while (c.charAt(0) == ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(name) == 0) {
			return c.substring(name.length, c.length);
		}
		}
		return "";
	}
`

func (q Script) UserData(name user.Data) qlova.String {
	q.Require(GetCookie)
	return q.wrap(`getCookie("` + string(name) + `")`)
}

const SetCookie = `
	function setCookie(cname, cvalue, exdays) {
		var d = new Date();
		d.setTime(d.getTime() + (exdays*24*60*60*1000));
		var expires = "expires="+ d.toUTCString();
		if (production) {
			document.cookie = cname + "=" + cvalue + ";" + expires + ";secure;path=/";
		} else {
			document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
		}
	}
`

func (q Script) SetUserData(name user.Data, value qlova.String) {
	q.Require(SetCookie)
	q.Javascript(`setCookie("` + string(name) + `", ` + raw(value) + `, 365);`)
}

func ToJavascript(f func(q Script), ctx ...internal.Context) []byte {
	if f == nil {
		return nil
	}

	var context internal.Context
	if len(ctx) > 0 {
		context = ctx[0]
	}

	return toJavascript(f, context)
}

func toJavascript(f func(q Script), context internal.Context) []byte {
	var program = qlova.Program(func(q qlova.Script) {
		var s = Script{&script{
			Script:  q,
			Context: context,
		}}
		s.js.q = s
		s.Time.Script = s
		//s.Go.Script = s
		f(s)
	})

	source := program.SourceCode(Javascript.Implementation{})
	if source.Error {
		panic(source.ErrorMessage)
	}

	return source.Data
}

func (q Script) JS() js {
	return q.js
}

func (q Script) Javascript(js string) {
	q.Raw("Javascript", language.Statement(js))
}

type Element struct {
	query string
	q     Script
}

func (q Script) Query(query qlova.String) Element {
	return Element{query: raw(query), q: q}
}

func (element Element) Run(method string) {
	element.q.Raw("Javascript", language.Statement(`document.querySelector(`+element.query+`).`+method+`();`))
}

type ExportedFunction struct {
	f reflect.Value
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

func (q Script) Run(f Function, args ...qlova.Type) {
	//.call(f, args...)
	q.Javascript(string(f) + "();")
}

//Export a Go function to Javascript. Don't use this for non-local apps! TODO enforce this
func (q Script) Call(f interface{}, args ...qlova.Type) qlova.Value {
	return q.call(f, args...)
}

func Handler(w http.ResponseWriter, r *http.Request, call string) {

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

type Unit qlova.String

func (unit Unit) Raw() string {
	return raw(qlova.String(unit))
}

func (script Script) Unit(unit complex128) Unit {
	return Unit(script.StringFromLanguageType(Javascript.String{
		Expression: language.Statement(css.Decode(unit)),
	}))
}

const SetClipboard = `
	const setClipboard = str => {
		const el = document.createElement('textarea');
		el.value = str;
		el.setAttribute('readonly', '');
		el.style.position = 'absolute';
		el.style.left = '-9999px';
		document.body.appendChild(el);
		const selected =
			document.getSelection().rangeCount > 0 ? document.getSelection().getRangeAt(0) : false;
		el.select();
		document.execCommand('copy');
		document.body.removeChild(el);
		if (selected) {
			document.getSelection().removeAllRanges();
			document.getSelection().addRange(selected);
		}
	};
`

func (script Script) SetClipboard(text String) {
	script.Require(SetClipboard)
	script.js.Run(`setClipboard`, text)
}
