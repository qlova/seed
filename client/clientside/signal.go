package clientside

import (
	"fmt"
	"io/ioutil"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/js"
)

//Signal is a communication tool.
type Signal struct {
	Name string

	address Address
	Memory  Memory

	Value string
}

//GetDefaultValue implements Variable
func (s *Signal) GetDefaultValue() client.Value {
	return js.Null()
}

//Variable implements Variable
func (s *Signal) Variable() (Address, Memory) {
	if s.address == "" {
		if s.Name != "" {
			s.address = Address(s.Name)
		} else {
			s.address = NewAddress()
		}
	}
	return s.address, s.Memory
}

//GetBool implements client.Bool
func (s *Signal) GetBool() js.Bool {
	return s.GetFunction().GetBool()
}

//GetValue implements client.Value
func (s *Signal) GetValue() js.Value {
	return s.GetFunction().Value
}

//GetFunction implements client.Function
func (s *Signal) GetFunction() js.Function {
	address, _ := s.Variable()
	return js.Func(fmt.Sprintf(`seed.variable.onchange["%v"]`, address))
}

//GetScript implements client.Script
func (s *Signal) GetScript() js.Script {
	address, _ := s.Variable()
	return js.Func(fmt.Sprintf(`seed.variable.onchange["%v"]`, address)).Run()
}

//On runs the given script when this signal is triggered.
func (s *Signal) On(do ...client.Script) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		js.NewCtx(ioutil.Discard, c)(client.NewScript(do...))

		var data data
		c.Read(&data)
		data.hooks = append(data.hooks, hook{
			variable: s,
			do:       client.NewScript(do...),
		})
		c.Write(data)
	})

}
