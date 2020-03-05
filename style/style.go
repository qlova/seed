package style

import (
	"fmt"
	"image/color"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
)

type data map[string]string

var seeds = make(map[int]data)

func convertColor(c color.Color) string {
	var r, g, b, a = c.RGBA()
	if a != 255 {
		c := fmt.Sprint("rgba(", (float64(r)/65535)*255, ",", (float64(g)/65535)*255, ",", (float64(b)/65535)*255, ",", float64(a)/65535, ")")
		return c
	} else {
		return fmt.Sprint("rgb(", r, ",", g, ",", b, ")")
	}
}

//SetTextColor sets the color of the seed.
func SetTextColor(c color.Color) seed.Option {
	return css.Set("color", convertColor(c))
}
