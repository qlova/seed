package script

import qlova "github.com/qlova/script"

//Object is a dynamic-content object.
type Object struct {
	Q Ctx
	qlova.Native
}

//NewObject creates a new Object.
func (q Ctx) NewObject() Object {
	return Object{
		Q:      q,
		Native: q.Value("{}").Native(),
	}
}

//Insert sets the value of the object at the given key.
func (o Object) Insert(key String, value qlova.Type) {
	var q = o.Q
	q.Javascript(o.Native.LanguageType().Raw() +
		`[` + key.LanguageType().Raw() + `] = ` +
		value.LanguageType().Raw() + ";")
}

//Lookup gets the value of the object at the given key.
func (o Object) Lookup(key String) Dynamic {
	var q = o.Q
	return q.Value(o.Native.LanguageType().Raw() +
		`[` + key.LanguageType().Raw() + `]`).Dynamic()
}

//Var calls Native.Var(...string).
func (o Object) Var(name ...string) Object {
	var variable = o.Native.Var(name...)
	return Object{
		Q:      o.Q,
		Native: variable,
	}
}
