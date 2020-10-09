package view

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientrender"
	"qlova.org/seed/js"
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

		c.With(client.OnLoad(js.Script(func(q js.Ctx) {
			fmt.Fprintf(q, `seed.view.ready(%v, "%v");`,
				client.Seed{c, q}.Element(), Name(starting))
		})))

		c.With(clientrender.On(js.Script(func(q js.Ctx) {

			fmt.Fprintf(q, `if (%[1]v.CurrentView) { %[1]v.CurrentView.args = %[2]v; if (%[1]v.CurrentView.onviewenter) %[1]v.CurrentView.onviewenter();  }`,
				client.Seed{c, q}.Element(), args.GetObject().String())
		})))
	})
}
