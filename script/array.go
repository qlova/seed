package script

import qlova "github.com/qlova/script"

//Array is a dynamic-content array.
type Array struct {
	Q Ctx
	qlova.Native
}

//NewArray creates a new Array.
func (q Ctx) NewArray() Array {
	return Array{
		Q:      q,
		Native: q.Value("[]").Native(),
	}
}

//Push pushes a new value to the array.
func (a Array) Push(v qlova.Type) {
	var q = a.Q
	q.Javascript(a.Native.LanguageType().Raw() + `.push(` + v.LanguageType().Raw() + ");")
}

//Index returns the value at the given index in the array.
func (a Array) Index(i Int) Dynamic {
	var q = a.Q
	return q.Value(a.Native.LanguageType().Raw() +
		`[` + i.LanguageType().Raw() + "]").Dynamic()
}

//Mutate sets the value at the given index in the array.
func (a Array) Mutate(i Int, v qlova.Type) {
	var q = a.Q
	q.Javascript(a.Native.LanguageType().Raw() +
		`[` + i.LanguageType().Raw() + "] = " +
		v.LanguageType().Raw() + ";")
}

//Var calls Native.Var(...string).
func (a Array) Var(name ...string) Array {
	var variable = a.Native.Var(name...)
	return Array{
		Q:      a.Q,
		Native: variable,
	}
}
