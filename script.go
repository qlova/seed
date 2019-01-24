package seed

import qlova "github.com/qlova/script"
import "github.com/qlova/seed/script"

type Script = script.Script

func NewVariable() script.Variable {
	return script.NewVariable()
}

//Return a scriptable version of this seed.
func (seed Seed) Script(q Script) script.Seed {
	return script.Seed{
		ID: seed.id,
		Q: q,
	}
}

//Set the text content of the seed.
func (seed Seed) SyncText(text *string) {
	var wrapper = func() string {
		return *text
	}
	
	seed.OnReady(func(q Script) {
		q.Javascript(`setInterval(function() {`)
			seed.Script(q).SetText(q.Call(wrapper).(qlova.String))
		q.Javascript(`}, 100)`)
	})
}