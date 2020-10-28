package text

import "qlova.org/seed/new/text/rich"

//Italic is shorthand for rich.Text(t).Italic()
func Italic(t rich.Text) rich.Text {
	return t.Italic()
}
