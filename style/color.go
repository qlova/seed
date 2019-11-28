package style

import (
	"bytes"
	"fmt"
	"image/color"
	"math"

	"github.com/qlova/seed/style/css"
)

//SetColor sets the color of this element.
func (style Style) SetColor(color color.Color) {
	style.CSS().SetBackgroundColor(css.Colour(color))
}

//SetTextColor sets the text color for this element.
func (style Style) SetTextColor(color color.Color) {
	style.Style.SetColor(css.Colour(color))
}

//Gradient is a
type Gradient struct {
	From, To  color.Color
	Colors    []color.Color
	Ratios    []float64
	Direction complex128

	Repeating, Radial, Circle bool

	Size    complex128
	Closest bool
	Side    bool
}

//CSS returns the gradient as a CSS function.
func (gradient Gradient) CSS() []byte {
	var buffer bytes.Buffer

	if len(gradient.Colors) == 0 {
		gradient.Colors = []color.Color{gradient.From, gradient.To}
	}

	if gradient.Repeating || len(gradient.Ratios) > 0 {
		buffer.WriteString(`repeating-`)
	}

	if gradient.Radial {

		buffer.WriteString(`radial-gradient(`)

		if gradient.Circle {
			buffer.WriteString(`circle `)
		}

		if gradient.Closest {
			buffer.WriteString(`closest-`)
		} else {
			buffer.WriteString(`farthest-`)
		}
		if gradient.Side {
			buffer.WriteString(`side `)
		} else {
			buffer.WriteString(`corner `)
		}

		if gradient.Size != 0 {
			fmt.Fprintf(&buffer, `at %v %v`, real(gradient.Size), imag(gradient.Size))
		}

		buffer.WriteByte(',')
	} else {
		buffer.WriteString(`linear-gradient(`)
		fmt.Fprintf(&buffer, "%vrad,", math.Atan2(imag(gradient.Direction), real(gradient.Direction))+math.Pi/2)
	}

	buffer.WriteString(css.Colour(gradient.Colors[0]).String())

	for i, col := range gradient.Colors {
		buffer.WriteByte(',')
		buffer.WriteString(css.Colour(col).String())
		if len(gradient.Ratios) > i {
			fmt.Fprintf(&buffer, " %v%%", gradient.Ratios[i])
		}
	}

	buffer.WriteByte(')')
	buffer.WriteByte(';')

	return buffer.Bytes()

}

//SetGradient sets the color of this element to be a gradient moving in direction from start color to end color.
func (style Style) SetGradient(gradient Gradient) {
	style.CSS().Set("background-image", string(gradient.CSS()))
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
