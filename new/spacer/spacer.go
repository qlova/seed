package spacer

import (
	"qlova.org/seed"
	"qlova.org/seed/web/css"
	"qlova.org/seed/set"
	"qlova.org/seed/web/css/units"

	"qlova.org/seed/new/html/div"
)

//New returns a new spacer that takes up the given amount of space.
func New(spacing units.Unit, options ...seed.Option) seed.Seed {
	return div.New(
		css.SetFlexBasis(css.Measure(spacing)),
		css.SetDisplay(css.Flex),
		css.SetFlexShrink(css.Zero),

		seed.Options(options),
	)
}

//Set the spacing of this seed.
func Set(spacing units.Unit) set.Style {
	return css.SetFlexBasis(css.Measure(spacing))
}
