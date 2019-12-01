package script

import (
	"fmt"

	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"

	Javascript "github.com/qlova/script/language/javascript"
)

//A nice interface to the Javascript world.
type js struct {
	q Ctx
}

//JS return the JS interface of script.
func (q Ctx) JS() js {
	return q.js
}

//Javascript inserts raw js into the script.
func (q Ctx) Javascript(js string, args ...interface{}) {
	var converted = make([]interface{}, len(args))
	for i := range args {
		if T, ok := args[i].(Type); ok {
			converted[i] = T.LanguageType().Raw()
		} else {
			converted[i] = args[i]
		}
	}

	if len(args) > 0 {
		q.Raw("Javascript", language.Statement(fmt.Sprintf(js, converted...)))
	} else {
		q.Raw("Javascript", language.Statement(fmt.Sprint(js)))
	}
}

//Value is any script value.
type Value struct {
	q   Ctx
	raw string
}

//Int returns the value as an Int.
func (v Value) Int() Int {
	return v.q.Script.ValueFromLanguageType(Javascript.Integer{Expression: language.Statement(v.raw)}).Int()
}

//String returns the value as a String.
func (v Value) String() String {
	return v.q.Script.ValueFromLanguageType(Javascript.String{Expression: language.Statement(v.raw)}).String()
}

//Bool returns the value as a bool.
func (v Value) Bool() qlova.Bool {
	return v.q.Script.ValueFromLanguageType(Javascript.Bit{Expression: language.Statement(v.raw)}).Bool()
}

//Float returns the value as a float.
func (v Value) Float() qlova.Float {
	return v.q.Script.ValueFromLanguageType(Javascript.Real{Expression: language.Statement(v.raw)}).Float()
}

//Native returns the value as a native value.
func (v Value) Native() qlova.Native {
	return v.q.Script.NativeFromLanguageType(Javascript.Native{Expression: language.Statement(v.raw)})
}

//Dynamic returns the value as a native value.
func (v Value) Dynamic() Dynamic {
	return Dynamic{
		Native: v.Native(),
		Q:      v.q,
	}
}

//File returns the value as an file.
func (v Value) File() File {
	return File{
		Native: v.Native(),
		Q:      v.q,
	}
}

//Location returns the value as a GeoLocation.
func (v Value) Location() Location {
	return Location{
		Native: v.Native(),
		Q:      v.q,
	}
}

//Array returns the value as an array.
func (v Value) Array() Array {
	return Array{
		Native: v.Native(),
		Q:      v.q,
	}
}

//Object returns the value as an object.
func (v Value) Object() Object {
	return Object{
		Native: v.Native(),
		Q:      v.q,
	}
}

//Promise returns the value as a promise.
func (v Value) Promise() Promise {
	return Promise{
		Native: v.Native().Var(),
		q:      v.q,
	}
}

//Unit returns the value as a unit.
func (v Value) Unit() Unit {
	return Unit(v.String())
}

//Value wraps a JS string as a value that can be cast to script.Type.
func (q Ctx) Value(format string, args ...interface{}) Value {

	var converted = make([]interface{}, len(args))
	for i := range args {
		if T, ok := args[i].(Type); ok {
			converted[i] = T.LanguageType().Raw()
		} else {
			converted[i] = args[i]
		}
	}

	if len(args) > 0 {
		return Value{q, fmt.Sprintf(format, converted...)}
	}
	return Value{q, format}
}

func (j js) Run(function string, args ...qlova.Type) {

	var converted string
	for i, arg := range args {
		converted += string(arg.LanguageType().Raw())
		if i <= len(args) {
			converted += ","
		}
	}

	j.q.Javascript(function + "(" + converted + ");")
}

func (j js) Call(function string, args ...qlova.Type) Value {

	var converted string
	for i, arg := range args {
		converted += string(arg.LanguageType().Raw())
		converted += ","
		if i <= len(args) {
		}
	}

	return Value{j.q, function + "(" + converted + ")"}
}
