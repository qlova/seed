package state

import (
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

//Bool is a global Boolean.
type Bool struct {
	Value
}

//NewBool returns a reference to a new global boolean.
func NewBool(options ...Option) Bool {
	return Bool{newValue("false", options...)}
}

//GetBool implements script.AnyBool
func (b Bool) GetBool() script.Bool {
	return b.get()
}

//GetValue implements script.AnyValue
func (b Bool) GetValue() script.Value {
	return b.get().Value
}

//Get the script.Bool for the global.Bool
func (b Bool) get() script.Bool {
	if b.raw != "" {
		return js.Bool{b.Value.getter()}
	}
	return js.Bool{js.NewValue(`(%v == "true")`, b.Value.get())}
}

//Set the global.Bool to be script.Bool
func (b Bool) set(q script.Ctx, value script.Bool) {
	b.Value.set(q, js.String{Value: js.NewValue(`(%v? "true": "")`, value)})
}

//Set allows setting the value of a String in the given script ctx.
func (b Bool) Set(value js.AnyBool) script.Script {
	return func(q script.Ctx) {
		b.set(q, value.GetBool())
	}
}

func (b Bool) If(options ...seed.Option) seed.Option {
	return State{Bool: b}.If(options...)
}

//Or returns a Bool that is true when either are true.
func (b Bool) Or(or js.AnyBool) Bool {
	var v = newValue("false")
	v.dependencies = &[]Value{b.Value}
	if other, ok := or.(AnyValue); ok {
		v.dependencies = &[]Value{b.Value, other.value()}
	}
	v.raw = fmt.Sprintf("(%v || %v)", b.GetBool().String(), or.GetBool().String())
	return Bool{v}
}

//And returns a Bool that is true when both are true.
func (b Bool) And(or js.AnyBool) Bool {
	var v = newValue("false")
	v.dependencies = &[]Value{b.Value}
	if other, ok := or.(AnyValue); ok {
		v.dependencies = &[]Value{b.Value, other.value()}
	}
	v.raw = fmt.Sprintf("(%v && %v)", b.GetBool().String(), or.GetBool().String())
	return Bool{v}
}
