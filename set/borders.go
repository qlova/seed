package set

import (
	"fmt"
	"image/color"

	"qlova.org/seed/use/css"
	"qlova.org/seed/use/css/units"
)

//BorderStyle determines how a border is rendered.
type BorderStyle int

//BorderStyles
const (
	Solid BorderStyle = iota
	Dashed
)

//Border sets the border style.
func Border(s BorderStyle) Style {
	switch s {
	case Solid:
		return css.SetBorderStyle(css.Solid)
	case Dashed:
		return css.SetBorderStyle(css.Dashed)
	default:
		panic("invalid border style")
	}
}

//Borderless removes the border of the seed.
func Borderless() Style {
	return css.Set("border", "none")
}

//BorderColor sets the border color of the seed.
func BorderColor(c color.Color) Style {
	return css.SetBorderColor(css.RGB{Color: c})
}

//BorderWidth sets the width of the border.
func BorderWidth(w units.Unit) Style {
	return css.Set("border-width", string(css.Measure(w).Rule()))
}

//Rounded sets how round this seed is, by rounding it's corners.
//If more than one argument is provided different corners are rounded differently.
//eg.
//
// 		Rounded(top, bottom)
// 		Rounded(top, bottom_left, bottom_right)
// 		Rounded(top_left, top_right, bottom_left, bottom_right)
//
// if more than 4 arguments are given, all but the first four are ignored.
func Rounded(first units.Unit, more ...units.Unit) css.Rule {
	switch len(more) {
	case 0:
		return css.Set("border-radius", string(css.Measure(first).Rule()))
	case 1:
		x := string(css.Measure(first).Rule())
		y := string(css.Measure(more[0]).Rule())
		return css.Set("border-radius", fmt.Sprintf(`%v %v`, y, x))
	case 2:
		top := string(css.Measure(first).Rule())
		bottomleft := string(css.Measure(more[0]).Rule())
		bottomright := string(css.Measure(more[1]).Rule())
		return css.Set("border-radius", fmt.Sprintf(`%v %v %v %v`, top, top, bottomright, bottomleft))
	default:
		topleft := string(css.Measure(first).Rule())
		topright := string(css.Measure(more[0]).Rule())
		bottomleft := string(css.Measure(more[1]).Rule())
		bottomright := string(css.Measure(more[2]).Rule())
		return css.Set("border-radius", fmt.Sprintf(`%v %v %v %v`, topleft, topright, bottomright, bottomleft))
	}
}

//TopLeftRounding sets the top-left rounding of this seed.
func TopLeftRounding(u units.Unit) Style {
	return css.Set("border-top-left-radius", css.Measure(u).String())
}

//TopRightRounding sets the top-right rounding of this seed.
func TopRightRounding(u units.Unit) Style {
	return css.Set("border-top-right-radius", css.Measure(u).String())
}

//BottomLeftRounding sets the bottom-left rounding of this seed.
func BottomLeftRounding(u units.Unit) Style {
	return css.Set("border-bottom-left-radius", css.Measure(u).String())
}

//BottomRightRounding sets the bottom-right rounding of this seed.
func BottomRightRounding(u units.Unit) Style {
	return css.Set("border-bottom-right-radius", css.Measure(u).String())
}
