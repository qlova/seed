package set

import (
	"fmt"

	"qlova.org/seed/css"
	"qlova.org/seed/units"
)

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
		x := string(css.Measure(first).Rule())
		top := string(css.Measure(more[0]).Rule())
		bottom := string(css.Measure(more[1]).Rule())
		return css.Set("border-radius", fmt.Sprintf(`%v %v %v`, top, x, bottom))
	default:
		left := string(css.Measure(first).Rule())
		right := string(css.Measure(more[0]).Rule())
		top := string(css.Measure(more[1]).Rule())
		bottom := string(css.Measure(more[2]).Rule())
		return css.Set("border-radius", fmt.Sprintf(`%v %v %v %v`, top, right, bottom, left))
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
