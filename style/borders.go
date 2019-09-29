package style

import (
	"fmt"
	"image/color"

	"github.com/qlova/seed/style/css"
)

//SetBorderless removes the border from this element.
func (style Style) SetBorderless() {
	style.SetBorderLeftWidth(css.Zero)
	style.SetBorderRightWidth(css.Zero)

	style.SetBorderTopWidth(css.Zero)
	style.SetBorderBottomWidth(css.Zero)
}

//SetBorder sets the border of this element to the specified color and thickness.
func (style Style) SetBorder(color color.Color, thickness int) {
	style.Set("border-left-width", fmt.Sprint(thickness, "px"))
	style.Set("border-right-width", fmt.Sprint(thickness, "px"))
	style.Set("border-top-width", fmt.Sprint(thickness, "px"))
	style.Set("border-bottom-width", fmt.Sprint(thickness, "px"))
	style.SetBorderColor(css.Colour(color))
}

//SetRoundedCorners sets this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCorners(radius Unit) {
	var value = css.Decode(radius)

	style.SetBorderBottomLeftRadius(value)
	style.SetBorderBottomRightRadius(value)

	style.SetBorderTopRightRadius(value)
	style.SetBorderTopLeftRadius(value)
}

//SetRoundedCornersBottom sets this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCornersBottom(radius Unit) {
	var value = css.Decode(radius)

	style.SetBorderBottomLeftRadius(value)
	style.SetBorderBottomRightRadius(value)
}

//SetRoundedCornersTop sets this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCornersTop(radius Unit) {
	var value = css.Decode(radius)

	style.SetBorderTopLeftRadius(value)
	style.SetBorderTopRightRadius(value)
}

//RemoveRoundedCorners removes any rounded corner specification for this element.
func (style Style) RemoveRoundedCorners() {

	style.SetBorderBottomLeftRadius(css.Unset)
	style.SetBorderBottomRightRadius(css.Unset)

	style.SetBorderTopRightRadius(css.Unset)
	style.SetBorderTopLeftRadius(css.Unset)
}
