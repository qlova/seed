package text

import (
	"image/color"

	"qlova.org/seed/new/text/rich"
)

//Italic is shorthand for rich.Text(t).Italic()
func Italic(t rich.Text) rich.Text {
	return t.Italic()
}

//Bold is shorthand for rich.Text(t).Bold()
func Bold(t rich.Text) rich.Text {
	return t.Bold()
}

//In is shorthand for rich.Text(t).In(c)
func In(c color.Color, t rich.Text) rich.Text {
	return t.In(c)
}
