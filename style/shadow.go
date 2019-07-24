package style

import (
	"fmt"
	"image/color"

	"github.com/qlova/seed/style/css"
)

//Shadow Defines a shadow that should be applied to the Element, with offset X and Y, Blur and of the specified color.
type Shadow struct {
	X, Y, Blur, Spread Unit
	Color              color.Color
}

//SetShadow sets the elements style to match the Shadow defintion.
func (style Style) SetShadow(shadow Shadow) {
	style.Set("box-shadow", fmt.Sprint(css.Decode(shadow.X), " ", css.Decode(shadow.Y), " ", css.Decode(shadow.Blur), " ", css.Decode(shadow.Spread), " ", css.Colour(shadow.Color)))
}
