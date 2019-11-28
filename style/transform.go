package style

import (
	"fmt"

	"github.com/qlova/seed/style/css"
)

//update updates the transform css of the style.
func (style *Style) update() {
	var transform = css.Rotate(0)
	transform = ""
	var changed bool

	if style.angle != nil {
		transform += css.Rotate(*style.angle)
		changed = true
	} else {
		transform += "rotate(var(--angle, 0))"
	}

	if style.rx != nil {
		transform += css.RotateX(*style.rx)
		changed = true
	} else {
		transform += "rotateX(var(--rx, 0))"
	}

	if style.y != nil && style.x != nil {

		transform += css.Translate(css.Decode(*style.x), css.Decode(*style.y))
		changed = true

	} else {
		if style.y != nil {
			transform += css.TranslateY(css.Decode(*style.y))
			changed = true
			transform += "translateX(var(--x, 0))"
		} else if style.x != nil {
			transform += css.TranslateX(css.Decode(*style.x))
			changed = true
			transform += "translateY(var(--y, 0))"
		} else {
			transform += "translate(var(--x, 0), var(--y, 0))"
		}
	}

	if style.scale != nil {
		transform += css.Scale(*style.scale, *style.scale)
		changed = true
	} else {
		transform += "scale(var(--scale, 1))"
	}

	if changed {
		style.CSS().SetTransform(transform)
	}
}

//Rotate the element by the given angle.
//This overrrides any previous calls to Angle.
func (style *Style) Rotate(angle float64) {
	style.CSS().Set("--angle", fmt.Sprint(angle, "deg"))
	style.angle = &angle
	style.update()
}

//RotateX the element by the given angle.
//This overrrides any previous calls to Angle.
func (style *Style) RotateX(angle float64) {
	style.CSS().Set("--rx", fmt.Sprint(angle, "deg"))
	style.rx = &angle
	style.update()
}

//Scale the element by the given scale.
//This overrrides any previous calls to Scale.
func (style *Style) Scale(scale float64) {
	style.CSS().Set("--scale", fmt.Sprint(scale))
	style.scale = &scale
	style.update()
}

//Translate the element by the given x and y values.
//This overrrides any previous calls to Translate.
func (style *Style) Translate(x, y complex128) {
	style.CSS().Set("--x", css.Decode(x).String())
	style.CSS().Set("--y", css.Decode(y).String())
	style.x = &x
	style.y = &y
	style.update()
}
