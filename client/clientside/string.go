package clientside

import (
	"github.com/google/uuid"
	"qlova.org/seed/client"
	"qlova.org/seed/js"
)

//String is a string variable in client memory.
type String struct {
	Name string

	address Address
	Memory  Memory

	Value string
}

//Variable implements Variable
func (s *String) Variable() (Address, Memory) {
	if s.address == "" {
		if s.Name != "" {
			s.address = Address(s.Name)
		} else {
			id, _ := uuid.NewRandom()
			s.address = Address(id.String())
		}
	}
	return s.address, s.Memory
}

//GetString implements client.String
func (s *String) GetString() js.String {
	address, memory := s.Variable()
	return js.String{Value: js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))}
}

//GetValue implements client.Value
func (s *String) GetValue() js.Value {
	return s.GetString().Value
}

//GetBool implements client.Value
func (s *String) GetBool() js.Bool {
	return s.GetValue().GetBool()
}

//GetDefaultValue implements Variable
func (s *String) GetDefaultValue() client.Value {
	return client.NewString(s.Value)
}

//Set returns a script that sets the string to the given literal.
func (s *String) Set(literal string) client.Script {
	return s.SetTo(client.NewString(literal))
}

//SetTo returns a script that sets the string to the given client string.
func (s *String) SetTo(value client.String) client.Script {
	address, memory := s.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), value)
}
