package script

import (
	"github.com/qlova/seed/internal"
	"github.com/qlova/seed/style/css"

	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"

	Javascript "github.com/qlova/script/language/javascript"
)

//Ctx is a script context. Providing access to script behaviours.
type Ctx struct {
	*ctx
}

type ctx struct {
	internal.Context
	qlova.Script

	js   js
	Time time
}

//Require inserts the provided dependency string in the head of the document.
func (q Ctx) Require(dependency string) {

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
func (q Ctx) RawString(s qlova.String) string {
	return raw(s)
}

func raw(s qlova.String) string {
	return string(s.LanguageType().(Javascript.String).Expression)
}

func (q Ctx) wrap(s string) qlova.String {
	return q.StringFromLanguageType(Javascript.String{
		Expression: language.Statement(s),
	})
}

//ToJavascript returns the given script encoded as Javascript.
func ToJavascript(f func(q Ctx), ctx ...internal.Context) []byte {
	if f == nil {
		return nil
	}

	var context internal.Context
	if len(ctx) > 0 {
		context = ctx[0]
	} else {
		context = internal.NewContext()
	}

	return toJavascript(f, context)
}

func toJavascript(f func(q Ctx), context internal.Context) []byte {
	var program = qlova.Program(func(q qlova.Script) {
		var s = Ctx{&ctx{
			Script:  q,
			Context: context,
		}}
		s.js.q = s
		s.Time.Ctx = s
		//s.Go.Script = s
		f(s)
	})

	source := program.SourceCode(Javascript.Implementation{})
	if source.Error {
		panic(source.ErrorMessage)
	}

	return source.Data
}

//Run runs a Javascript function with the given arguments.
func (q Ctx) Run(f Function, args ...qlova.Type) {
	q.Javascript(string(f) + "();")
}

//Unit is a display unit, eg. px, em, %
type Unit qlova.String

//Raw returns the unit as a raw string.
func (unit Unit) Raw() string {
	return raw(qlova.String(unit))
}

//Unit returns a script.Unit from the given unit.
func (q Ctx) Unit(unit complex128) Unit {
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
func (q Ctx) SetClipboard(text String) {
	q.Require(SetClipboard)
	q.js.Run(`setClipboard`, text)
}
