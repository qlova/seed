package clientside

import (
	"time"

	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//Duration is a time.Duration in client.Memory.
type Duration struct {
	MemoryAddress

	Value time.Duration
}

//GetDuration implements client.Time
func (s *Duration) GetDuration() js.Value {
	address, memory := s.Variable()
	return js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))
}

//GetValue implements client.Value
func (s *Duration) GetValue() js.Value {
	return s.GetDuration()
}

//GetBool implements client.Value
func (s *Duration) GetBool() js.Bool {
	return s.GetValue().GetBool()
}

//GetDefaultValue implements Variable
func (s *Duration) GetDefaultValue() client.Value {
	return client.NewDuration(s.Value)
}

//Set returns a script that sets the string to the given literal.
func (s *Duration) Set(literal time.Duration) client.Script {
	return s.SetTo(client.NewDuration(literal))
}

//SetTo returns a script that sets the string to the given client string.
func (s *Duration) SetTo(value client.Duration) client.Script {
	address, memory := s.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), value)
}

//Time is a time.Time value in client memory.
type Time struct {
	MemoryAddress

	Value time.Time
}

//GetTime implements client.Time
func (s *Time) GetTime() js.Value {
	address, memory := s.Variable()
	return js.Call(js.Func("q.getvar"),
		client.NewString(string(address)), client.NewString(string(memory)))
}

//GetValue implements client.Value
func (s *Time) GetValue() js.Value {
	return s.GetTime()
}

//GetBool implements client.Value
func (s *Time) GetBool() js.Bool {
	return s.GetValue().GetBool()
}

//GetDefaultValue implements Variable
func (s *Time) GetDefaultValue() client.Value {
	return client.NewTime(s.Value)
}

//Set returns a script that sets the string to the given literal.
func (s *Time) Set(literal time.Time) client.Script {
	return s.SetTo(client.NewTime(literal))
}

//SetTo returns a script that sets the string to the given client string.
func (s *Time) SetTo(value client.Time) client.Script {
	address, memory := s.Variable()
	return js.Run(js.Func("q.setvar"), client.NewString(string(address)),
		client.NewString(string(memory)), value)
}
