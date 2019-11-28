package style

import "github.com/qlova/seed/style/css"

//Clip clips the content if it would go out of it's container.
func (style Style) Clip() {
	style.CSS().SetOverflow(css.Hidden)
}

//SetScrollable sets that this can be scrolled vertically.
func (style Style) SetScrollable() {
	style.CSS().SetOverflowY(css.Auto)
	style.CSS().SetOverflowX(css.Hidden)
	style.Compress()
	style.CSS().Set("-webkit-overflow-scrolling", "touch")
	style.CSS().Set("-webkit-overscroll-behavior", "contain")
	style.CSS().Set("overscroll-behavior", "contain")
}

//SetFullyScrollable sets that this can be scrolled horizontally or vertically.
func (style Style) SetFullyScrollable() {
	style.CSS().SetOverflow(css.Auto)
	style.CSS().Set("-webkit-overflow-scrolling", "touch")
	style.CSS().Set("-webkit-overscroll-behavior", "contain")
	style.CSS().Set("overscroll-behavior", "contain")
}

//SetNotScrollable sets that this cannot be scrolled.
func (style Style) SetNotScrollable() {
	style.CSS().SetOverflow(css.Hidden)
	style.CSS().SetOverflowX(css.Hidden)
	style.CSS().SetOverflowY(css.Hidden)
}
