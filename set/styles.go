package set

import (
	"image/color"

	"qlova.org/seed"
	"qlova.org/seed/use/css"
)

//Style is a setter that wraps css styling.
type Style interface {
	seed.Option
	Rules() css.Rules
}

//Scrollable sets this to seed to be scrollable if it's height overflows.
//If the container is getting cut off, ensure that a parent seed as MinHeight set to 0.
func Scrollable() Style {
	return css.Rules{
		css.SetOverflowY(css.Auto),
		css.SetOverflowX(css.Hidden),
		css.Set("-webkit-overflow-scrolling", "touch"),
		css.Set("-webkit-overscroll-behavior", "contain"),
		css.Set("overscroll-behavior", "contain"),
	}
}

//Color sets the color of this seed.
func Color(c color.Color) Style {
	return css.SetBackgroundColor(css.RGB{Color: c})
}

//Clipped sets this seed to clip any children that cross the border.
func Clipped() Style {
	return css.SetOverflow(css.Hidden)
}

//Hidden removes this seed from taking space and being visible.
func Hidden() Style {
	return css.SetDisplay(css.None)
}

//Visible sets the seed to be visible.
func Visible() css.Rule {
	return css.SetDisplay(css.Flex)
}
