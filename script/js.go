package script

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

//A nice interface to the Javascript world.
type js struct {
	q Script
}

type value struct {
	q Script
	raw string
}

func (v value) String() qlova.String {
	return v.q.Script.ValueFromLanguageType(Javascript.String{Expression:language.Statement(v.raw)}).String()
}

func (v value) Bool() qlova.Bool {
	return v.q.Script.ValueFromLanguageType(Javascript.Bit{Expression:language.Statement(v.raw)}).Bool()
}

func (j js) Run(function string, args ...qlova.Type) {
	
	var converted string
	for i, arg := range args {
		converted += string(arg.LanguageType().Raw())
		if i <= len(args) {
			converted += ","
		}
	}
	
	j.q.Javascript(function+"("+converted+");")
}

func (j js) Call(function string, args ...qlova.Type) value {
	
	var converted string
	for i, arg := range args {
		converted += string(arg.LanguageType().Raw())
			converted += ","
		if i <= len(args) {
		}
	}
	
	return value{j.q, function+"("+converted+")"}
}
