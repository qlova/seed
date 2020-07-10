package set

import (
	"fmt"

	"qlova.org/seed/css"
	"qlova.org/seed/units"
)

//Padding sets the inner-spacing of this seed.
//If more than one argument is provided different sides are padded.
//eg.
//
// 		Padding(x, y)
// 		Padding(x, top, bottom)
// 		Padding(left, right, top, bottom)
//
// if more than 4 arguments are given, all but the first four are ignored.
func Padding(first units.Unit, more ...units.Unit) css.Rule {
	switch len(more) {
	case 0:
		return css.Set("padding", string(css.Measure(first).Rule()))
	case 1:
		x := string(css.Measure(first).Rule())
		y := string(css.Measure(more[0]).Rule())
		return css.Set("padding", fmt.Sprintf(`%v %v`, y, x))
	case 2:
		x := string(css.Measure(first).Rule())
		top := string(css.Measure(more[0]).Rule())
		bottom := string(css.Measure(more[1]).Rule())
		return css.Set("padding", fmt.Sprintf(`%v %v %v`, top, x, bottom))
	default:
		left := string(css.Measure(first).Rule())
		right := string(css.Measure(more[0]).Rule())
		top := string(css.Measure(more[1]).Rule())
		bottom := string(css.Measure(more[2]).Rule())
		return css.Set("padding", fmt.Sprintf(`%v %v %v %v`, top, right, bottom, left))
	}
}

//PaddingTop sets the top inner-spacing of this seed.
func PaddingTop(u units.Unit) Style {
	return css.SetPaddingTop(css.Measure(u))
}

//PaddingBottom sets the bottom inner-spacing of this seed.
func PaddingBottom(u units.Unit) Style {
	return css.SetPaddingBottom(css.Measure(u))
}

//PaddingLeft sets the inner-spacing to the left of this seed.
func PaddingLeft(u units.Unit) Style {
	return css.SetPaddingLeft(css.Measure(u))
}

//PaddingRight sets the inner-spacing to the right of this seed.
func PaddingRight(u units.Unit) Style {
	return css.SetPaddingRight(css.Measure(u))
}

//Margin sets the outer-spacing of this seed.
//If more than one argument is provided different sides are padded.
//eg.
//
// 		Margin(x, y)
// 		Margin(x, top, bottom)
// 		Margin(left, right, top, bottom)
//
// if more than 4 arguments are given, all but the first four are ignored.
func Margin(first units.Unit, more ...units.Unit) css.Rule {
	switch len(more) {
	case 0:
		return css.Set("margin", string(css.Measure(first).Rule()))
	case 1:
		x := string(css.Measure(first).Rule())
		y := string(css.Measure(more[0]).Rule())
		return css.Set("margin", fmt.Sprintf(`%v %v`, y, x))
	case 2:
		x := string(css.Measure(first).Rule())
		top := string(css.Measure(more[0]).Rule())
		bottom := string(css.Measure(more[1]).Rule())
		return css.Set("margin", fmt.Sprintf(`%v %v %v`, top, x, bottom))
	default:
		left := string(css.Measure(first).Rule())
		right := string(css.Measure(more[0]).Rule())
		top := string(css.Measure(more[1]).Rule())
		bottom := string(css.Measure(more[2]).Rule())
		return css.Set("margin", fmt.Sprintf(`%v %v %v %v`, top, right, bottom, left))
	}
}

//MarginTop sets the top outer-spacing of this seed.
func MarginTop(u units.Unit) Style {
	return css.SetMarginTop(css.Measure(u))
}

//MarginBottom sets the bottom outer-spacing of this seed.
func MarginBottom(u units.Unit) Style {
	return css.SetMarginBottom(css.Measure(u))
}

//MarginLeft sets the outer-spacing to the left of this seed.
func MarginLeft(u units.Unit) Style {
	return css.SetMarginLeft(css.Measure(u))
}

//MarginRight sets the outer-spacing to the right of this seed.
func MarginRight(u units.Unit) Style {
	return css.SetMarginRight(css.Measure(u))
}
