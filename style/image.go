package style

import "github.com/qlova/seed/style/css"

//SetContain makes sure that this maintains its aspect ratio.
func (style Style) SetContain() {
	style.CSS().SetObjectFit(css.Contain)
}

//Stretch causes the image to stretch.
func (style Style) Stretch() {
	style.Style.SetObjectFit(css.Cover)
}
