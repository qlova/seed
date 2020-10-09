package icon

import (
	"qlova.org/seed/asset"
	"qlova.org/seed/s/text/rich"
)

//Inline an icon inside text.
func Inline(src string) rich.Text {
	src = asset.Path(src)
	return rich.Icon(src)
}
