package clientside

import (
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//Float64 is an float64 variable in client memory.
type Float64 struct {
	MemoryAddress

	Value float64
}

//GetNumber implements client.Number
func (v *Float64) GetNumber() js.Number {
	address, memory := v.Variable()
	return js.Number{Value: js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))}
}

//GetValue implements client.Value
func (v *Float64) GetValue() js.Value {
	return v.GetNumber().Value
}

//GetBool implements client.Value
func (v *Float64) GetBool() js.Bool {
	return v.GetValue().GetBool()
}

//GetDefaultValue implements Variable
func (v *Float64) GetDefaultValue() client.Value {
	return client.NewFloat64(v.Value)
}

//Set returns a script that sets the int to the given literal.
func (v *Float64) Set(literal float64) client.Script {
	return v.SetTo(js.NewNumber(literal))
}

//SetTo returns a script that sets the string to the given client string.
func (v *Float64) SetTo(value client.Float) client.Script {
	address, memory := v.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), value)
}
