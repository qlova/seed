package set

import (
	"fmt"
	"image/color"

	"qlova.org/seed"
	"qlova.org/seed/use/css"
	"qlova.org/seed/use/css/units"
)

//Shadow Defines a shadow that should be applied to the Element, with offset X and Y, Blur and of the specified color.
type Shadow struct {
	X, Y, Blur, Spread units.Unit
	Color              color.Color

	Inset bool
}

func (shadow Shadow) AddTo(c seed.Seed) {
	var inset = ""
	if shadow.Inset {
		inset = "inset "
	}

	if shadow.Color == nil {
		shadow.Color = color.NRGBA{0, 0, 0, 255}
	}

	css.Set("box-shadow",
		fmt.Sprint(
			inset,
			css.Measure(shadow.X).String(), " ",
			css.Measure(shadow.Y).String(), " ",
			css.Measure(shadow.Blur).String(), " ",
			css.Measure(shadow.Spread).String(), " ",
			css.RGB{Color: shadow.Color}.Rule(),
		),
	).AddTo(c)
}
