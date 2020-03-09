package style

import (
	"fmt"
	"image/color"

	"github.com/qlova/seed"

	"github.com/qlova/seed/css"
)

type Unit complex64

func (u Unit) String() string {
	if u == 0 {
		return "0"
	}
	return fmt.Sprintf("%v%%", real(u))
}

type Style interface {
	seed.Option
	Rules() css.Rules
}

func convertColor(c color.Color) string {
	var r, g, b, a = c.RGBA()
	if a != 255 {
		c := fmt.Sprint("rgba(", (float64(r)/65535)*255, ",", (float64(g)/65535)*255, ",", (float64(b)/65535)*255, ",", float64(a)/65535, ")")
		return c
	} else {
		return fmt.Sprint("rgb(", r, ",", g, ",", b, ")")
	}
}

//Translate the element by the given x and y values.
//This overrrides any previous calls to Translate.
func Translate(x, y Unit) css.Rules {
	return css.Rules{
		css.Set("--x", x.String()),
		css.Set("--y", y.String()),
	}
}

//SetTextColor sets the color of the seed.
func SetTextColor(c color.Color) css.Rule {
	return css.Set("color", convertColor(c))
}

//SetColumn sets the seed to behave as a column.
func SetColumn() css.Rule {
	return css.SetFlexDirection(css.Column)
}

//SetHidden removes the seed.
func SetHidden() css.Rule {
	return css.SetDisplay(css.None)
}

//SetVisible sets the seed to be visible.
func SetVisible() css.Rule {
	return css.SetDisplay(css.Flex)
}
