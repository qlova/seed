package state

import (
	"fmt"

	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

type AnyFloat interface{}

//Float is a global Float.
type Float struct {
	Value
}

//NewFloat returns a reference to a new global float.
func NewFloat(initial float64) Float {
	return Float{newValue(fmt.Sprintf("%#v", initial))}
}

//GetNumber implements script.AnyNumber
func (i Float) GetNumber() script.Number {
	return i.get()
}

//GetValue implements script.AnyValue
func (i Float) GetValue() script.Value {
	return i.get().Value
}

//GetValue implements script.AnyValue
func (i Float) GetBool() script.Bool {
	return i.GetValue().GetBool()
}

func (i Float) get() script.Number {
	return js.Number{Value: js.NewValue(`parseFloat(%v)`, i.Value.get())}
}

//Set allows setting the value of a String in the given script ctx.
func (f Float) Set(value script.Number) script.Script {
	return func(q script.Ctx) {
		f.set(q, value)
	}
}

//SetL sets the value of the Int with a literal.
func (i Float) SetL(value float64) script.Script {
	return func(q script.Ctx) {
		i.set(q, q.Number(value))
	}
}

func (i Float) set(q script.Ctx, value script.Number) {
	i.Value.set(q, js.String{Value: js.NewValue(`(%v).toString()`, value)})
}
