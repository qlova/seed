package seed

import "github.com/qlova/seed/script"

type Bool = script.Bool
type Script = script.Script

var functions = make(map[string]func(Script))

//Define a new function that can be called from any Script context.
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

//Return a scriptable version of this seed.
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

//Set the text content of the seed.
/*func (seed Seed) SyncText(text *string) {
	var wrapper = func() string {
		return *text
	}

	seed.OnReady(func(q Script) {
		q.Javascript(`setInterval(function() {`)
		seed.Script(q).SetText(q.Call(wrapper).String())
		q.Javascript(`}, 100)`)
	})
}*/
