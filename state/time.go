package state

import (
	"fmt"

	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

type AnyTime interface{}

//Time is a global Time.
type Time struct {
	Value
}

//NewTime returns a reference to a new global Time.
func NewTime() Time {
	return Time{newValue(fmt.Sprintf("0"))}
}

//GetNumber implements script.AnyNumber
func (i Time) GetNumber() script.Number {
	return i.get()
}

//GetValue implements script.AnyValue
func (i Time) GetValue() script.Value {
	return i.get().Value
}

//GetValue implements script.AnyValue
func (i Time) GetBool() script.Bool {
	return i.GetValue().GetBool()
}

func (i Time) get() script.Number {
	return js.Number{Value: js.NewValue(`parseFloat(%v)`, i.Value.get())}
}

//Set allows setting the value of a String in the given script ctx.
func (f Time) SetNow() script.Script {
	return func(q script.Ctx) {
		f.set(q, js.Number{js.NewValue(`Date.now()`)})
	}
}

func (i Time) set(q script.Ctx, value script.Number) {
	i.Value.set(q, js.String{Value: js.NewValue(`(%v).toString()`, value)})
}
