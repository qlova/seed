package seed

import "github.com/qlova/seed/script"

type (
	//Bool aliases to the script.Bool type.
	Bool = script.Bool

	//Args aliases to the script.Args type.
	Args = script.Args
)

var functions = make(map[string]func(script.Ctx))

//NewFunction defines a new function that can be called from any Ctx context.
func NewFunction(f func(script.Ctx), names ...string) script.Function {
	var name string
	if len(names) > 0 {
		name = names[0]
	} else {
		name = script.Unique()
	}

	//TODO auto dependencies.
	functions[name] = f

	return script.Function(name)
}

//Ctx returns the script context of this seed.
func (seed Seed) Ctx(q script.Ctx) script.Seed {
	if seed.Template {
		return script.Seed{
			Native: seed.id,
			Q:      q,
		}
	}
	return script.Seed{
		ID: seed.id,
		Q:  q,
	}
}
