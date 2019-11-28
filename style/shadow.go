package style

import (
	"fmt"
	"image/color"

	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
)

//Shadow Defines a shadow that should be applied to the Element, with offset X and Y, Blur and of the specified color.
type Shadow struct {
	X, Y, Blur, Spread unit.Unit
	Color              color.Color
}

//SetShadow sets the elements style to match the Shadow defintion.
func (style Style) SetShadow(shadow Shadow) {
	style.CSS().Set("box-shadow", fmt.Sprint(css.Decode(shadow.X), " ", css.Decode(shadow.Y), " ", css.Decode(shadow.Blur), " ", css.Decode(shadow.Spread), " ", css.Colour(shadow.Color)))
}

//RemoveShadow removes the shadow on this element.
func (style Style) RemoveShadow() {
	style.CSS().Set("box-shadow", "unset")
}

//SetTextShadow sets the elements text style to match the Shadow defintion.
func (style Style) SetTextShadow(shadow Shadow) {
	style.CSS().Set("text-shadow", fmt.Sprint(css.Decode(shadow.X), " ", css.Decode(shadow.Y), " ", css.Decode(shadow.Blur), " ", css.Colour(shadow.Color)))
}

//RemoveTextShadow removes the text shadow on this element.
func (style Style) RemoveTextShadow() {
	style.CSS().Set("text-shadow", "")
}
