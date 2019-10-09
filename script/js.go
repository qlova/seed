package script

import (
	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"

	Javascript "github.com/qlova/script/language/javascript"
)

//A nice interface to the Javascript world.
type js struct {
	q Ctx
}

type value struct {
	q   Ctx
	raw string
}

func (v value) String() qlova.String {
	return v.q.Script.ValueFromLanguageType(Javascript.String{Expression: language.Statement(v.raw)}).String()
}

func (v value) Bool() qlova.Bool {
	return v.q.Script.ValueFromLanguageType(Javascript.Bit{Expression: language.Statement(v.raw)}).Bool()
}

func (v value) Float() qlova.Float {
	return v.q.Script.ValueFromLanguageType(Javascript.Real{Expression: language.Statement(v.raw)}).Float()
}

func (v value) Native() qlova.Native {
	return v.q.Script.NativeFromLanguageType(Javascript.Native{Expression: language.Statement(v.raw)})
}

func (v value) Promise() Promise {
	return Promise{
		Native: v.Native().Var(),
		q:      v.q,
	}
}

func (v value) Unit() Unit {
	return Unit(v.String())
}

//Value wraps a JS string as a value that can be cast to script.Type.
func (q Ctx) Value(raw string) value {
	return value{q, raw}
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

func (j js) Call(function string, args ...qlova.Type) value {

	var converted string
	for i, arg := range args {
		converted += string(arg.LanguageType().Raw())
		converted += ","
		if i <= len(args) {
		}
	}

	return value{j.q, function + "(" + converted + ")"}
}
