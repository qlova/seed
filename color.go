package seed

import (
	"encoding/hex"
	"image/color"
)

//Some colors as constants.
const (
	White Hex = "#ffffff"
	Black Hex = "#000000"
	Red   Hex = "#ff0000"
	Green Hex = "#00ff00"
	Blue  Hex = "#0000ff"
)

//RGB returns a new color from the specified r, g and b values.
func RGB(r, g, b uint8) color.Color {
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

//RGBA returns a new color from the specified r, g, b and a values.
func RGBA(r, g, b, a uint8) color.Color {
	return color.RGBA{R: r, G: g, B: b, A: a}
}

//Hex is a color in hex representation ie #ffffff
//strings of this format can be cast directly to the Hex type.
//ie Hex("#ffffff")
type Hex string

//RGBA returns premultiplied r, g, b & a values for this hex color.
func (h Hex) RGBA() (r, g, b, a uint32) {
	var c [4]byte
	c[3] = 255
	hex.Decode(c[:], []byte(h[1:]))
	r, g, b, a = uint32(c[0]), uint32(c[1]), uint32(c[2]), uint32(c[3])
	if a != 255 {
		r = uint32((float64(r) / 255) * 65535)
		g = uint32((float64(g) / 255) * 65535)
		b = uint32((float64(b) / 255) * 65535)
		a = uint32((float64(a) / 255) * 65535)
	}
	return
}
