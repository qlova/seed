package client

import (
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

//Memory is a type of client memory for SideValues.
type Memory string

//Memory types.
const (
	ShortTermMemory Memory = "short"
	SessionMemory   Memory = "session"
	LongTermMemory  Memory = "storage"

	LocalMemory Memory = "local"
)

//Value is a readonly client-typed value.
type Value js.AnyValue

//SideValue is a client-side variable value.
type SideValue struct {
	Value

	Memory
	Key string
}

//GetValue implements Value.
func (v SideValue) GetValue() js.Value {
	return js.NewValue("q").Call("GetValue",
		NewString(string(v.Memory)), NewString(v.Key))
}

//Set sets the value of the SideValue.
func (v SideValue) Set(value Value) Script {
	return script.New(
		js.NewValue("q").Run("SetValue",
			NewString(string(v.Memory)), NewString(v.Key), value),
	)
}
