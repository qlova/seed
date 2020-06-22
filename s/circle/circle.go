package circle

import (
	"qlova.org/seed"
	"qlova.org/seed/s/column"
	"qlova.org/seed/style"
)

//New returns a new circle.
func New(radius style.Unit, options ...seed.Option) seed.Seed {
	return column.New(
		style.SetSize(radius, radius),
		style.SetRoundedCorners(50),

		seed.Options(options),
	)
}
