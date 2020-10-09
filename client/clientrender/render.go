package clientrender

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/js"
)

//On executes the provided script whenever this seed is rendered.
func On(do client.Script) seed.Option {
	return client.On("render", do.GetScript())
}

//If renders the provided options if the condition is true.
func If(condition client.Bool, options ...seed.Option) seed.Option {

	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(condition, c)

		c.With(On(client.If(condition, js.Script(func(q js.Ctx) {
			for _, option := range options {
				if option == nil {
					continue
				}
				if other, ok := option.(seed.Seed); ok {
					client.Seed{other, q}.AddTo(client.Seed{c, q})
				} else {
					option.AddTo(client.Seed{c, q})
				}
			}
		})).Else(js.Script(func(q js.Ctx) {
			for _, option := range options {
				if option == nil {
					continue
				}
				if other, ok := option.(seed.Seed); ok {
					client.Seed{c, q}.Undo(client.Seed{other, q})
				} else {
					client.Seed{c, q}.Undo(option)
				}
			}
		}))))
	})
}
