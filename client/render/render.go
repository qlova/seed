package render

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/script"
)

//On executes the provided script whenever this seed is rendered.
func On(do client.Script) seed.Option {
	return script.On("render", do)
}

//If renders the provided options if the condition is true.
func If(condition client.Bool, options ...seed.Option) seed.Option {

	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(condition, c)

		c.With(On(client.If(condition, func(q script.Ctx) {
			for _, option := range options {
				if option == nil {
					continue
				}
				if other, ok := option.(seed.Seed); ok {
					script.Scope(other, q).AddTo(script.Scope(c, q))
				} else {
					option.AddTo(script.Scope(c, q))
				}
			}
		}).Else(func(q script.Ctx) {
			for _, option := range options {
				if option == nil {
					continue
				}
				if other, ok := option.(seed.Seed); ok {
					script.Scope(c, q).Undo(script.Scope(other, q))
				} else {
					script.Scope(c, q).Undo(option)
				}
			}
		})))
	})
}
