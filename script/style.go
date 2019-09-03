package script

import (
	"fmt"
	"image/color"
	"math"

	qlova "github.com/qlova/script"
	"github.com/qlova/seed/style/css"
)

type Color struct {
	String
}

func (q Script) Color(c color.Color) Color {
	r, g, b, a := c.RGBA()
	return Color{q.String(fmt.Sprint("rgba(", r, ",", g, ",", b, ",", a, ")"))}
}

func (q Script) Hex(s string) Color {
	return Color{q.String(s)}
}

func (seed Seed) Hidden() qlova.Bool {
	return seed.Q.Value(`(getComputedStyle(` + seed.Element() + `, null).display == "none")`).Bool()
}

//SetColor sets the color of this seed.
func (seed Seed) SetColor(c Color) {
	seed.Set("background-color", `"+`+c.LanguageType().Raw()+`+"`)
}

//SetInvisible causes the seed to still take up space but be hidden from view.
func (seed Seed) SetInvisible() {
	seed.Set("visibility", css.Hidden.String())
}

//SetVisible causes the seed to take up space and be visible.
func (seed Seed) SetVisible() {
	seed.Set("visibility", css.Visible.String())
	seed.Set("display", css.Flex.String())
}

//SetGradient sets the gradient of the seed.
func (seed Seed) SetGradient(direction complex128, start, end Color) {
	seed.Set("background-image", fmt.Sprint("linear-gradient(", math.Atan2(imag(direction), real(direction))+math.Pi/2, `deg,"+`, css.ColorValue(start.LanguageType().Raw()), `+","+`, css.ColorValue(end.LanguageType().Raw()), `+")`))
}

//SetGradient sets the gradient of the seed.
func (seed Seed) ClearGradient() {
	seed.Set("background-image", "")
}
