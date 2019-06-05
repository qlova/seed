package seed

import "image/color"
import "encoding/hex"

const White Hex = "#ffffff"

func RGB(r, g, b uint8) color.Color {
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func RGBA(r, g, b, a uint8) color.Color {
	return color.RGBA{R: r, G: g, B: b, A: a}
}

type Hex string

func (h Hex) RGBA() (r, g, b, a uint32) {
	var c [4]byte
	c[3] = 255
	hex.Decode(c[:], []byte(h[1:]))
	return uint32(c[0]), uint32(c[1]), uint32(c[2]), uint32(c[3])
}
