package script

import (
	"fmt"
	"image/color"
	"math"

	qlova "github.com/qlova/script"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/css"
)

//Color is a script interface to a color value.
type Color struct {
	String
}

//Color returns a script color based on the color.Color.
func (q Ctx) Color(c color.Color) Color {
	r, g, b, a := c.RGBA()
	return Color{q.String(fmt.Sprint("rgba(", r, ",", g, ",", b, ",", a, ")"))}
}

//Hex returns a color based on the hex value.
func (q Ctx) Hex(s string) Color {
	return Color{q.String(s)}
}

//Hidden returns true if the seed is hidden.
func (seed Seed) Hidden() qlova.Bool {
	return seed.Q.Value(`(getComputedStyle(` + seed.Element() + `, null).display == "none")`).Bool()
}

//Visible returns true if the seed is visible.
func (seed Seed) Visible() qlova.Bool {
	return seed.Q.Value(`(getComputedStyle(` + seed.Element() + `, null).display != "none")`).Bool()
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

//Column returns true if this is a column layout.
func (seed Seed) Column() Bool {
	return seed.Q.Value(`(getComputedStyle(` + seed.Element() + `, null).flexDirection == "column")`).Bool()
}

//SetColumn causes the seed to act as a column layout.
func (seed Seed) SetColumn() {
	seed.Set("flex-direction", css.Column.String())
}

//Row returns true if this is a row layout.
func (seed Seed) Row() Bool {
	return seed.Q.Value(`(getComputedStyle(` + seed.Element() + `, null).flexDirection == "row")`).Bool()
}

//SetRow causes the seed to act as a row layout.
func (seed Seed) SetRow() {
	seed.Set("flex-direction", css.Row.String())
}

//SetGradient sets the gradient of the seed.
func (seed Seed) SetGradient(direction complex128, start, end Color) {
	seed.Set("background-image", fmt.Sprint("linear-gradient(", math.Atan2(imag(direction), real(direction))+math.Pi/2, `deg,"+`, css.ColorValue(start.LanguageType().Raw()), `+","+`, css.ColorValue(end.LanguageType().Raw()), `+")`))
}

//ClearGradient clears the gradient of the seed.
func (seed Seed) ClearGradient() {
	seed.Set("background-image", "")
}

//SetTint sets the tint of this seed. Color must be known at runtime.
func (seed Seed) SetTint(tint color.Color) {
	var s = style.New()
	s.SetTint(tint)
	var f = s.Style.Get("filter")
	seed.Set("filter", f[:len(f)-1])
}

//RemoveTint removes the tint of this seed. Color must be known at runtime.
func (seed Seed) RemoveTint() {
	seed.Set("filter", "unset")
}

//Translate sets the transform of this seed to the specified translation.
func (seed Seed) Translate(x, y Unit) {
	seed.Javascript(seed.Element() + `.style.setProperty("--x", "` + x.Raw() + `");`)
	seed.Javascript(seed.Element() + `.style.setProperty("--y", "` + y.Raw() + `");`)
	seed.Set("transform", "rotate(var(--angle, 0)) scale(var(--scale, 1)) translate(var(--x, 0), var(--y, 0))")
}

//Scale sets the scale of this seed to the specified scalar.
func (seed Seed) Scale(scalar Float) {
	seed.Javascript(seed.Element() + `.style.setProperty("--scale", ` + scalar.LanguageType().Raw() + `);`)
	seed.Set("transform", "rotate(var(--angle, 0)) scale(var(--scale, 1)) translate(var(--x, 0), var(--y, 0))")
}
