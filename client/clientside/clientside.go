package clientside

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

type hook struct {
	variable Variable
	do       client.Script
	render   seed.Set
}

type data struct {
	seed.Data

	hooks []hook
}

//Render rerenders the given seed as a client script.
func Render(c seed.Seed) client.Script {
	return js.Func("await c.r").Run(script.Q, js.NewString(script.ID(c)))
}

//Compound values are scanned by Hook
type Compound interface {
	Components() []client.Value
}

//Hook renders the given seed whenever the value changes.
//if v is not a Variable or Compound, this is a noop
func Hook(v client.Value, c seed.Seed) {
	variable, ok := v.(Variable)
	if !ok {
		compound, ok := v.(Compound)
		if !ok {
			return
		}

		components := compound.Components()

		for _, component := range components {
			Hook(component, c)
		}

		return
	}
	var data data
	c.Read(&data)
	data.hooks = append(data.hooks, hook{
		variable: variable,
		render:   seed.NewSet(c),
	})
	c.Write(data)
}
