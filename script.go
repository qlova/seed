package seed

import "github.com/qlova/seed/script"

type (
	//Bool aliases to the script.Bool type.
	Bool = script.Bool

	//Script aliases to the script.Bool type.
	Script = script.Script

	//Args aliases to the script.Args type.
	Args = script.Args
)

var functions = make(map[string]func(Script))

//NewFunction defines a new function that can be called from any Script context.
func NewFunction(f func(Script), names ...string) script.Function {
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

//Script returns a scriptable version of this seed.
func (seed Seed) Script(q Script) script.Seed {
	if seed.template {
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
