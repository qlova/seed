package spacer

import (
	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/style"

	"qlova.org/seed/s/html/div"
)

//New returns a new row.
func New(space style.Unit, options ...seed.Option) seed.Seed {
	return div.New(css.SetFlexBasis(space.Unit()), css.Set("display", "flex").And(options...))
}
