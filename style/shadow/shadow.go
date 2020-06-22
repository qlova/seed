package shadow

import (
	"fmt"
	"image/color"

	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/style"
)

//Shadow Defines a shadow that should be applied to the Element, with offset X and Y, Blur and of the specified color.
type Shadow struct {
	X, Y, Blur, Spread style.Unit
	Color              color.Color

	Inset bool
}

//New is an alias for Gradient.
type New = Shadow

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
			shadow.X.Unit().Rule(), " ",
			shadow.Y.Unit().Rule(), " ",
			shadow.Blur.Unit().Rule(), " ",
			shadow.Spread.Unit().Rule(), " ",
			css.RGB{Color: shadow.Color}.Rule(),
		),
	).AddTo(c)
}
