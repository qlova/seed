package view

import (
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

//Set adds and sets an initial view to the seed.
func Set(starting View) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		if starting == nil {
			return
		}

		ControllerOf(c).Goto(starting)

		c.Add(script.OnReady(func(q script.Ctx) {
			fmt.Fprintf(q, `seed.view.ready(%v, "%v");`,
				script.Scope(c, q).Element(), Name(starting))
		}))
	})
}
