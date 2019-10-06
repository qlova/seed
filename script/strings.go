package script

import (
	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"
	Javascript "github.com/qlova/script/language/javascript"
)

//SubString returns a slice of s, from start to end.
func (q Ctx) SubString(s String, start, end Int) String {
	return q.js.Call(s.LanguageType().Raw()+".substr", start, end).String()
}

//Contains returns true if text contains match.
func (q Ctx) Contains(text, match qlova.String) qlova.Bool {
	return q.Script.BoolFromLanguageType(Javascript.Bit{Expression: language.Statement(raw(text) + `.includes(` + raw(match) + `)`)})
}
