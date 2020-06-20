package margin

import (
	"fmt"

	"qlova.org/seed/css"
	"qlova.org/seed/style"
)

//Set sets the margin values, Set(all), Set(x, y), Set(left, right, top, bottom)
func Set(margin ...style.Unit) css.Rule {
	const (
		left = iota
		right
		top
		bottom
	)

	if len(margin) == 1 {
		return css.Set("margin", fmt.Sprint(margin[0].Unit().Rule()))
	}
	if len(margin) == 2 {
		return css.Set("margin", fmt.Sprint(margin[1].Unit().Rule(), " ", margin[0].Unit().Rule()))
	}
	if len(margin) == 4 {
		return css.Set("margin", fmt.Sprint(margin[top].Unit().Rule(),
			" ", margin[right].Unit().Rule(),
			" ", margin[bottom].Unit().Rule(),
			" ", margin[left].Unit().Rule()))
	}
	panic("invalid arguments")
}

//SetTop sets the top margin.
func SetTop(margin style.Unit) css.Rule {
	return css.SetMarginTop(margin.Unit())
}

//SetBottom sets the bottom margin.
func SetBottom(margin style.Unit) css.Rule {
	return css.SetMarginBottom(margin.Unit())
}

//SetRight sets the right margin.
func SetRight(margin style.Unit) css.Rule {
	return css.SetMarginRight(margin.Unit())
}

//SetLeft sets the letf margin.
func SetLeft(margin style.Unit) css.Rule {
	return css.SetMarginLeft(margin.Unit())
}
