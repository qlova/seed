package view

import (
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
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

		c.Add(script.OnReady(func(q script.Ctx) {
			fmt.Fprintf(q, `seed.view.ready(%v, "%v");`,
				script.Scope(c, q).Element(), Name(starting))
		}))

		c.Add(state.OnRefresh(func(q script.Ctx) {
			fmt.Fprintf(q, `if (%[1]v.CurrentView) { %[1]v.CurrentView.args = %[2]v; %[1]v.CurrentView.rerender(); }`,
				script.Scope(c, q).Element(), args.GetObject().String())
		}))
	})
}
