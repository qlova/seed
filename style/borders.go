package style

import (
	"fmt"
	"image/color"

	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
)

//SetBorderless removes the border from this element.
func (style Style) SetBorderless() {
	style.CSS().SetBorderLeftWidth(css.Zero)
	style.CSS().SetBorderRightWidth(css.Zero)

	style.CSS().SetBorderTopWidth(css.Zero)
	style.CSS().SetBorderBottomWidth(css.Zero)
}

//SetBorder sets the border of this element to the specified color and thickness.
func (style Style) SetBorder(color color.Color, thickness int) {
	style.CSS().Set("border-left-width", fmt.Sprint(thickness, "px"))
	style.CSS().Set("border-right-width", fmt.Sprint(thickness, "px"))
	style.CSS().Set("border-top-width", fmt.Sprint(thickness, "px"))
	style.CSS().Set("border-bottom-width", fmt.Sprint(thickness, "px"))
	style.CSS().SetBorderColor(css.Colour(color))
	style.CSS().SetBorderStyle(css.Solid)
}

//SetRoundedCorners sets this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCorners(radius unit.Unit) {
	var value = css.Decode(radius)

	style.CSS().SetBorderBottomLeftRadius(value)
	style.CSS().SetBorderBottomRightRadius(value)

	style.CSS().SetBorderTopRightRadius(value)
	style.CSS().SetBorderTopLeftRadius(value)
}

//SetRoundedCornersLeft sets this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCornersLeft(radius unit.Unit) {
	var value = css.Decode(radius)

	style.CSS().SetBorderBottomLeftRadius(value)
	style.CSS().SetBorderTopLeftRadius(value)
}

//SetRoundedCornersRight sets this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCornersRight(radius unit.Unit) {
	var value = css.Decode(radius)

	style.CSS().SetBorderBottomRightRadius(value)
	style.CSS().SetBorderTopRightRadius(value)
}

//SetRoundedCornersBottom sets this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCornersBottom(radius unit.Unit) {
	var value = css.Decode(radius)

	style.CSS().SetBorderBottomLeftRadius(value)
	style.CSS().SetBorderBottomRightRadius(value)
}

//SetRoundedCornersTop sets this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCornersTop(radius unit.Unit) {
	var value = css.Decode(radius)

	style.CSS().SetBorderTopLeftRadius(value)
	style.CSS().SetBorderTopRightRadius(value)
}

//RemoveRoundedCorners removes any rounded corner specification for this element.
func (style Style) RemoveRoundedCorners() {

	style.CSS().SetBorderBottomLeftRadius(css.Unset)
	style.CSS().SetBorderBottomRightRadius(css.Unset)

	style.CSS().SetBorderTopRightRadius(css.Unset)
	style.CSS().SetBorderTopLeftRadius(css.Unset)
}
