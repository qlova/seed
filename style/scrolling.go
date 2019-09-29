package style

import "github.com/qlova/seed/style/css"

//Clip clips the content if it would go out of it's container.
func (style Style) Clip() {
	style.SetOverflow(css.Hidden)
}

//SetScrollable sets that this can be scrolled vertically.
func (style Style) SetScrollable() {
	style.SetOverflowY(css.Auto)
	style.SetOverflowX(css.Hidden)
	style.Compress()
	style.Set("-webkit-overflow-scrolling", "touch")
	style.Set("-webkit-overscroll-behavior", "contain")
	style.Set("overscroll-behavior", "contain")
}

//SetFullyScrollable sets that this can be scrolled horizontally or vertically.
func (style Style) SetFullyScrollable() {
	style.SetOverflow(css.Auto)
	style.Set("-webkit-overflow-scrolling", "touch")
	style.Set("-webkit-overscroll-behavior", "contain")
	style.Set("overscroll-behavior", "contain")
}

//SetNotScrollable sets that this cannot be scrolled.
func (style Style) SetNotScrollable() {
	style.SetOverflow(css.Hidden)
	style.SetOverflowX(css.Hidden)
	style.SetOverflowY(css.Hidden)
}
