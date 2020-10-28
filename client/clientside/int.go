package clientside

import (
	"qlova.org/seed/client"
	"qlova.org/seed/web/js"
)

//Int is an int variable in client memory.
type Int struct {
	Name string

	address Address
	Memory  Memory

	Value int
}

//Variable implements Variable
func (i *Int) Variable() (Address, Memory) {
	if i.address == "" {
		if i.Name != "" {
			i.address = Address(i.Name)
		} else {
			i.address = NewAddress()
		}
	}
	return i.address, i.Memory
}

//GetNumber implements client.Int
func (i *Int) GetNumber() js.Number {
	address, memory := i.Variable()
	return js.Number{Value: js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))}
}

//GetValue implements client.Value
func (i *Int) GetValue() js.Value {
	return i.GetNumber().Value
}

//GetBool implements client.Value
func (i *Int) GetBool() js.Bool {
	return i.GetValue().GetBool()
}

//GetDefaultValue implements Variable
func (i *Int) GetDefaultValue() client.Value {
	return client.NewInt(i.Value)
}

//Set returns a script that sets the int to the given literal.
func (i *Int) Set(literal int) client.Script {
	address, memory := i.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), client.NewInt(literal))
}

//SetTo returns a script that sets the int to the given client.Int.
func (i *Int) SetTo(v client.Int) client.Script {
	address, memory := i.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), v)
}

//Add adds the given literal value to the Int.
func (i *Int) Add(literal int) client.Script {
	address, memory := i.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), i.GetNumber().Plus(client.NewInt(literal)))
}
