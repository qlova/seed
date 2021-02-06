package style

import (
	"qlova.org/seed"
	"qlova.org/seed/new/font"
	"qlova.org/seed/set"
	"qlova.org/seed/use/css/units/px"
	"qlova.tech/rgb"
)

var Font = font.New("fixedsys.ttf")
var Text = seed.Options{Font}

var Border = seed.Options{
	set.Border(set.Solid),
	set.BorderColor(rgb.Black),
	set.BorderWidth(px.New(3.0)),
}
