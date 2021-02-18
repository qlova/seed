package change

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/client/screen"
	"qlova.org/seed/set"
	"qlova.org/seed/use/js"
)

//On executes the provided script whenever this seed is rendered.
func On(do client.Script) seed.Option {
	return client.On("render", do.GetScript())
}

//When applies the given changes when the condition is true.
func When(condition client.Bool, changes ...seed.Option) seed.Option {
	if q, ok := condition.(screen.SizeQuery); ok {
		return seed.NewOption(func(c seed.Seed) {

			var stylingChanges []set.Style

			for _, child := range changes {
				if styling, ok := child.(set.Style); ok {
					stylingChanges = append(stylingChanges, styling)
				} else {
					panic("dynamic options depending on a screen.SizeQuery is not supported yet, please submit an issue")
				}
			}

			set.Query(q.Media(), stylingChanges...).AddTo(c)
		})
	}

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
