package clientside

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/web/js"
)

type hook struct {
	variable Variable
	do       client.Script
	render   seed.Set
}

type data struct {
	hooks []hook
}

//Render rerenders the given seed as a client script.
func Render(c seed.Seed) client.Script {
	return js.Func("await c.r").Run(js.NewValue("q"), js.NewString(client.ID(c)))
}

//Hook renders the given seed whenever the value changes.
//if v is not a Variable or Compound, this is a noop
func Hook(v client.Value, c seed.Seed) {
	variable, ok := v.(Variable)
	if !ok {
		compound, ok := v.(client.Compound)
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
	c.Load(&data)

	switch mode, q := client.Seed(c); mode {
	case client.AddTo:
		address, _ := variable.Variable()
		fmt.Fprintf(q, `seed.variable.hook['%v'].push('%v');`, address, client.ID(c))
	case client.Undo:
		fmt.Fprintf(q, `if (seed.variable.hook['%v']) seed.variable.hook['%v'].splice(seed.variable.hook['%v'].indexOf('%v'), 1);`, address, address, address, client.Element(c))
	default:
		data.hooks = append(data.hooks, hook{
			variable: variable,
			render:   seed.NewSet(c),
		})
	}

	c.Save(data)
}
