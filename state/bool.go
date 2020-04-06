package state

import (
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
	return js.Bool{js.NewValue(`(%v == "true")`, b.Value.get())}
}

//Set the global.Bool to be script.Bool
func (b Bool) set(q script.Ctx, value script.Bool) {
	b.Value.set(q, js.String{Value: js.NewValue(`(%v).toString()`, value)})
}
