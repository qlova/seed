package clientside

import (
	"fmt"
	"io/ioutil"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//Signal is a communication tool.
type Signal struct {
	MemoryAddress

	Value string
}

//GetDefaultValue implements Variable
func (s *Signal) GetDefaultValue() client.Value {
	return js.Null()
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
		c.Load(&data)
		data.hooks = append(data.hooks, hook{
			variable: s,
			do:       client.NewScript(do...),
		})
		c.Save(data)
	})

}
