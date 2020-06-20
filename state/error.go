package state

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

func Error(into String) seed.Option {
	return script.OnError(func(q script.Ctx, err script.Error) {
		q(into.Set(err.String))
	})
}
