package view

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client/render"
	"qlova.org/seed/script"
)

//Set adds and sets an initial view to the seed.
func Set(starting View) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		if starting == nil {
			return
		}

		//Sort out script arguments of the page.
		_, args := parseArgs(starting, c)

		ControllerOf(c).Goto(starting)

		c.With(script.OnReady(func(q script.Ctx) {
			fmt.Fprintf(q, `seed.view.ready(%v, "%v");`,
				script.Scope(c, q).Element(), Name(starting))
		}))

		c.With(render.On(func(q script.Ctx) {

			fmt.Fprintf(q, `if (%[1]v.CurrentView) { %[1]v.CurrentView.args = %[2]v; if (%[1]v.CurrentView.onviewenter) %[1]v.CurrentView.onviewenter();  }`,
				script.Scope(c, q).Element(), args.GetObject().String())
		}))
	})
}
