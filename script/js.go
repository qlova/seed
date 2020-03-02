package script

import (
	"fmt"

	"github.com/qlova/script"
	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"
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
		if v, ok := args[i].(script.Value); ok {
			converted[i] = q.Raw(v)
		} else {
			converted[i] = args[i]
		}
	}

	if len(args) > 0 {
		q.Write([]byte(fmt.Sprintf(js, converted...)))
	} else {
		q.Write([]byte(fmt.Sprint(js)))
	}
}

//Value is any script value.
type value struct {
	q   Ctx
	raw string
}

//Int returns the value as an Int.
func (v value) Int() Int {
	return Int{language.Expression(v.q, v.raw)}
}

//String returns the value as a String.
func (v value) String() String {
	return String{language.Expression(v.q, v.raw)}
}

//Bool returns the value as a bool.
func (v value) Bool() qlova.Bool {
	return Bool{language.Expression(v.q, v.raw)}
}

//Float returns the value as a float.
func (v value) Float() qlova.Float {
	return Float{language.Expression(v.q, v.raw)}
}

//Native returns the value as a native value.
func (v value) Native() qlova.Native {
	return Native{language.Expression(v.q, v.raw)}
}

//Dynamic returns the value as a native value.
func (v value) Dynamic() Dynamic {
	return Dynamic{
		Native: v.Native(),
		Q:      v.q,
	}
}

//File returns the value as an file.
func (v value) File() File {
	return File{
		Native: v.Native(),
		Q:      v.q,
	}
}

//Location returns the value as a GeoLocation.
func (v value) Location() Location {
	return Location{
		Native: v.Native(),
		Q:      v.q,
	}
}

//Array returns the value as an array.
func (v value) Array() Array {
	return Array{
		Native: v.Native(),
		Q:      v.q,
	}
}

//Object returns the value as an object.
func (v value) Object() Object {
	return Object{
		Native: v.Native(),
		Q:      v.q,
	}
}

//Promise returns the value as a promise.
func (v value) Promise() Promise {

	var n = v.Native()
	n.Var()

	return Promise{
		Native: n,
		q:      v.q,
	}
}

//Unit returns the value as a unit.
func (v value) Unit() Unit {
	return Unit(v.String())
}

//Value wraps a JS string as a value that can be cast to script.Type.
func (q Ctx) Value(format string, args ...interface{}) value {

	var converted = make([]interface{}, len(args))
	for i := range args {
		if v, ok := args[i].(script.Value); ok {
			converted[i] = q.Raw(v)
		} else {
			converted[i] = args[i]
		}
	}

	if len(args) > 0 {
		return value{q, fmt.Sprintf(format, converted...)}
	}
	return value{q, format}
}

func (j js) Run(function string, args ...qlova.Value) {

	var converted string
	for i, arg := range args {
		converted += j.q.Raw(arg)
		if i <= len(args) {
			converted += ","
		}
	}

	j.q.Javascript(function + "(" + converted + ");")
}

func (j js) Call(function string, args ...qlova.Value) value {

	var converted string
	for i, arg := range args {
		converted += j.q.Raw(arg)
		converted += ","
		if i <= len(args) {
		}
	}

	return value{j.q, function + "(" + converted + ")"}
}
