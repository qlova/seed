package set

import (
	"bytes"
	"fmt"
	"image/color"
	"math"

	"qlova.org/seed"
	"qlova.org/seed/use/css"
)

//Gradient is a color gradient from one color to another.
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

//AddTo implements seed.Option
func (gradient Gradient) AddTo(c seed.Seed) {
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

	buffer.WriteString(string(css.RGB{Color: gradient.Colors[0]}.Rule()))

	for i, col := range gradient.Colors {
		buffer.WriteByte(',')
		buffer.WriteString(string(css.RGB{Color: col}.Rule()))
		if len(gradient.Ratios) > i {
			fmt.Fprintf(&buffer, " %v%%", gradient.Ratios[i])
		}
	}

	buffer.WriteByte(')')
	buffer.WriteByte(';')

	css.Set("background", buffer.String()).AddTo(c)
}
