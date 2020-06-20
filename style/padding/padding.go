package padding

import (
	"fmt"

	"qlova.org/seed/css"
	"qlova.org/seed/style"
)

//Set sets the padding values.
func Set(padding ...style.Unit) css.Rule {
	if len(padding) == 1 {
		return css.Set("padding", fmt.Sprint(padding[0].Unit().Rule()))
	}
	if len(padding) == 2 {
		return css.Set("padding", fmt.Sprint(padding[1].Unit().Rule(), " ", padding[0].Unit().Rule()))
	}
	panic("invalid arguments")
}

//SetTop sets the top margin.
func SetTop(padding style.Unit) css.Rule {
	return css.SetPaddingTop(padding.Unit())
}

//SetBottom sets the bottom margin.
func SetBottom(padding style.Unit) css.Rule {
	return css.SetPaddingBottom(padding.Unit())
}

//SetRight sets the right margin.
func SetRight(padding style.Unit) css.Rule {
	return css.SetPaddingRight(padding.Unit())
}

//SetLeft sets the letf margin.
func SetLeft(padding style.Unit) css.Rule {
	return css.SetPaddingLeft(padding.Unit())
}
