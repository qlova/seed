package icon

import (
	"qlova.org/seed/assets"
	"qlova.org/seed/new/text/rich"
)

//Inline an icon inside text.
func Inline(src string) rich.Text {
	src = assets.Path(src)
	return rich.Icon(src)
}
