package change

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/web/js"
)

//On executes the provided script whenever this seed is rendered.
func On(do client.Script) seed.Option {
	return client.On("render", do.GetScript())
}

//When applies the given changes when the condition is true.
func When(condition client.Bool, changes ...seed.Option) seed.Option {

	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(condition, c)

		c.With(On(client.If(condition, js.Script(func(q js.Ctx) {
			for _, option := range changes {
				if option == nil {
					continue
				}
				if _, ok := option.(seed.Seed); ok {
					continue
				} else {
					client.Option(option, q).AddTo(c)
				}
			}
		})).Else(js.Script(func(q js.Ctx) {
			for _, option := range changes {
				if option == nil {
					continue
				}
				if _, ok := option.(seed.Seed); ok {
					continue
				} else {
					client.Reverse(option, q).AddTo(c)
				}
			}
		}))))
	})
}
