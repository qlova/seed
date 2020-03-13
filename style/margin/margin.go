package margin

import (
	"fmt"

	"github.com/qlova/seed/css"
	"github.com/qlova/seed/style"
)

//Set sets the margin values.
func Set(margin ...style.Unit) css.Rule {
	if len(margin) == 1 {
		return css.Set("margin", fmt.Sprint(margin[0].Unit().Rule()))
	}
	if len(margin) == 2 {
		return css.Set("margin", fmt.Sprint(margin[1].Unit().Rule(), " ", margin[0].Unit().Rule()))
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
