package spacer

import (
	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/units"

	"qlova.org/seed/s/html/div"
)

//New returns a new spacer that takes up the given amount of space.
func New(spacing units.Unit) seed.Seed {
	return div.New(
		css.SetFlexBasis(css.Measure(spacing)),
		css.SetDisplay(css.Flex),
		css.SetFlexShrink(css.Zero),
	)
}
