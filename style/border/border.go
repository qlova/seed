package border

import (
	"image/color"

	"qlova.org/seed/css"
	"qlova.org/seed/style"
)

//Remove removes the borders of this seed.
func Remove() css.Rule {
	return css.Set("border", "none")
}

func SetColor(c color.Color) css.Rule {
	return css.SetBorderColor(css.RGB{Color: c})
}

func SetWidth(width style.Unit) css.Rule {
	return css.SetBorderWidth(width.Unit())
}

func SetSolid() css.Rule {
	return css.SetBorderStyle(css.Solid)
}

func SetDashed() css.Rule {
	return css.SetBorderStyle(css.Dashed)
}