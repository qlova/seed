package expander

import (
	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/s/html/div"
)

//New returns a new expander that expands to fill empty space.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		css.SetFlexGrow(css.Number(1)),
	)
}

//Set sets the seed to expand to fill up space.
func Set() css.Rule {
	return css.SetFlexGrow(css.Number(1))
}
