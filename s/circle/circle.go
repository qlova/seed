package circle

import (
	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/s/column"
	"qlova.org/seed/set"
	"qlova.org/seed/units"
	"qlova.org/seed/units/percentage/of"
)

//New returns a new circle.
func New(options ...seed.Option) seed.Seed {
	return column.New(
		css.SetFlexShrink(css.Zero),
		seed.Options(options),
	)
}

//Set sets the radius of the circle.
func Set(radius units.Unit) seed.Option {
	return seed.Options{
		set.Size(radius, radius),
		set.Rounded(50 % of.Parent),
	}
}
