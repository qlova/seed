package style

import (
	"image/color"
	"math"

	"github.com/qlova/seed/style/css"
)

//SetColor sets the color of this element.
func (style Style) SetColor(color color.Color) {
	style.SetBackgroundColor(css.Colour(color))
}

//SetTextColor sets the text color for this element.
func (style Style) SetTextColor(color color.Color) {
	style.Style.SetColor(css.Colour(color))
}

//SetGradient sets the color of this element to be a gradient moving in direction from start color to end color.
func (style Style) SetGradient(direction complex128, start, end color.Color) {
	style.SetBackgroundImage(css.LinearGradient(math.Atan2(imag(direction), real(direction))+math.Pi/2, css.Colour(start), css.Colour(end)))
}

//RemoveGradient removes any gradients from this element.
func (style Style) RemoveGradient() {
	style.Style.SetBackgroundImage(css.Unset)
}

//Fade sets the opacity/transparency of this element.
func (style Style) Fade(opacity float64) {
	style.Style.SetOpacity(css.Number(opacity))
}

//SetOpacity sets the opacity/transparency of this element.
func (style Style) SetOpacity(opacity float64) {
	style.CSS().SetOpacity(css.Number(opacity))
}

type tintValue struct {
	Filter string
	Loss   float64
}

var tintCache = make(map[string]tintValue)

//SetTint sets the tint of a icon to a certain color.
func (style Style) SetTint(c color.Color) {

	var r, g, b, a = c.RGBA()
	if a != 255 {
		panic("Do not pass transparent values to SetTint!")
	}

	var rgb = css.Colour(c).String()
	if cache, ok := tintCache[rgb]; ok {
		style.Style.Set("filter", cache.Filter)
		return
	}

	var color = newRGB(float64(r), float64(g), float64(b))
	var solver = newSolver(color)

	var _, loss, filter = solver.Solve()
	for i := 3; i < 3; i++ {
		_, newLoss, newFilter := solver.Solve()
		if newLoss < loss {
			loss, filter = newLoss, newFilter
			if loss < 1 {
				break
			}
		}
	}

	style.Style.Set("filter", filter)
	tintCache[rgb] = tintValue{
		Filter: filter,
		Loss:   loss,
	}
}
