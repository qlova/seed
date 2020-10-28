package clientside

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//Value is an untyped variable in client memory.
type Value struct {
	Name string

	address Address
	Memory  Memory

	Value client.Value
}

//Variable implements Variable
func (v *Value) Variable() (Address, Memory) {
	if v.address == "" {
		if v.Name != "" {
			v.address = Address(v.Name)
		} else {
			v.address = NewAddress()
		}
	}
	return v.address, v.Memory
}

//GetBool implements client.Bool
func (v *Value) GetBool() js.Bool {
	return v.GetValue().GetBool()
}

//GetValue implements client.Value
func (v *Value) GetValue() js.Value {
	address, memory := v.Variable()

	return js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))
}

//GetDefaultValue implements Variable
func (v *Value) GetDefaultValue() client.Value {
	return js.Null()
}

//SetTo returns a script that sets the value to the given client.Value.
func (v *Value) SetTo(value client.Value) client.Script {
	address, memory := v.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), value)
}

//OnChange runs the given script when the value of this string is changed.
func (v *Value) OnChange(do ...client.Script) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Load(&data)
		data.hooks = append(data.hooks, hook{
			variable: v,
			do:       client.NewScript(do...),
		})
		c.Save(data)
	})

}
