package script

import "github.com/qlova/seed/internal"

import "github.com/qlova/seed/style/css"
import qlova "github.com/qlova/script"

import "github.com/qlova/script/language"
import Javascript "github.com/qlova/script/language/javascript"

//Script is an alias to Context.
type Script struct {
	*script
}

//Context is a script context. Providing access to script behaviours.
type Context = Script

type script struct {
	internal.Context
	qlova.Script

	js   js
	Time time
}

//Require inserts the provided dependency string in the head of the document.
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

//RawString is an internal function and should not be used.
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

//ToJavascript returns the given script encoded as Javascript.
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

//JS return the JS interface of script.
func (q Script) JS() js {
	return q.js
}

//Javascript inserts raw js into the script.
func (q Script) Javascript(js string) {
	q.Raw("Javascript", language.Statement(js))
}

//Run runs a Javascript function with the given arguments.
func (q Script) Run(f Function, args ...qlova.Type) {
	q.Javascript(string(f) + "();")
}

//Unit is a display unit, eg. px, em, %
type Unit qlova.String

//Raw returns the unit as a raw string.
func (unit Unit) Raw() string {
	return raw(qlova.String(unit))
}

//Unit returns a script.Unit from the given unit.
func (q Script) Unit(unit complex128) Unit {
	return Unit(q.StringFromLanguageType(Javascript.String{
		Expression: language.Statement(css.Decode(unit)),
	}))
}

//SetClipboard is the JS code requried for Clipboard support.
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

//SetClipboard sets the clipboard to the provided string.
func (q Script) SetClipboard(text String) {
	q.Require(SetClipboard)
	q.js.Run(`setClipboard`, text)
}
