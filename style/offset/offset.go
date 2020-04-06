package offset

import (
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/style"
)

//Set sets the offset values.
func Set(offset ...style.Unit) css.Rules {
	if len(offset) == 1 {
		return css.Rules{
			SetTop(offset[0]),
			SetBottom(offset[0]),
			SetRight(offset[0]),
			SetLeft(offset[0]),
		}
	}
	if len(offset) == 2 {
		return css.Rules{
			SetTop(offset[1]),
			SetBottom(offset[1]),
			SetRight(offset[0]),
			SetLeft(offset[0]),
		}
	}
	panic("invalid arguments")
}

//SetTop sets the top offset.
func SetTop(offset style.Unit) css.Rule {
	return css.SetTop(offset.Unit())
}

//SetBottom sets the bottom offset.
func SetBottom(offset style.Unit) css.Rule {
	return css.SetBottom(offset.Unit())
}

//SetRight sets the right offset.
func SetRight(offset style.Unit) css.Rule {
	return css.SetRight(offset.Unit())
}

//SetLeft sets the letf offset.
func SetLeft(offset style.Unit) css.Rule {
	return css.SetLeft(offset.Unit())
}
