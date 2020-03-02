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
func (o Object) Insert(key String, value qlova.Value) {
	var q = o.Q
	q.Javascript(q.Raw(o.Native) +
		`[` + q.Raw(key) + `] = ` +
		q.Raw(value.T()) + ";")
}

//Lookup gets the value of the object at the given key.
func (o Object) Lookup(key String) Dynamic {
	var q = o.Q
	return q.Value(q.Raw(o.Native) +
		`[` + q.Raw(key) + `]`).Dynamic()
}

//Var calls Native.Var(...string).
func (o Object) Var(name ...string) Object {
	var variable = o.Native
	variable.Var(name...)
	return Object{
		Q:      o.Q,
		Native: variable,
	}
}
