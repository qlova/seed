package set

import (
	"qlova.org/seed/css"
	"qlova.org/seed/units"
)

//Translation sets the translation of this seed.
func Translation(x, y units.Unit) css.Rules {
	return css.Rules{
		css.Set("--x", css.Measure(x).String()),
		css.Set("--y", css.Measure(y).String()),
		css.Set("transform", "translate(var(--x, 0), var(--y, 0)) rotate(var(--angle, 0)) scale(var(--scale, 1), var(--scale, 1))"),
	}
}
