package script

import (
	qlova "github.com/qlova/script"
)

type AnyString = qlova.AnyString

//SubString returns a slice of s, from start to end.
func (q Ctx) SubString(s String, start, end Int) String {
	return q.Value(`%v.substr(%v, %v)`, s, start, end).String()
}

//Contains returns true if text contains match.
func (q Ctx) Contains(text, match qlova.String) qlova.Bool {
	return q.Value(`%v.includes(%v)`, text, match).Bool()
}
