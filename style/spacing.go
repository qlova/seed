package style

import "github.com/qlova/seed/style/css"

//SetInnerSpacing sets the spacing within this element, takes em, vm, px or percentage values.
func (style Style) SetInnerSpacing(x, y Unit) {
	style.SetPaddingLeft(css.Decode(x))
	style.SetPaddingRight(css.Decode(x))

	style.SetPaddingTop(css.Decode(y))
	style.SetPaddingBottom(css.Decode(y))
}

//SetInnerSpacingTop sets the inner spacing top, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingTop(value Unit) {
	style.SetPaddingTop(css.Decode(value))
}

//SetInnerSpacingLeft sets the inner spacing left, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingLeft(value Unit) {
	style.SetPaddingLeft(css.Decode(value))
}

//SetInnerSpacingBottom sets the inner spacing bottom, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingBottom(value Unit) {
	style.SetPaddingBottom(css.Decode(value))
}

//SetInnerSpacingRight sets the inner spacing right, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingRight(value Unit) {
	style.SetPaddingRight(css.Decode(value))
}

//SetOuterSpacing sets the outer spacing of this element, takes em, vm, px or percentage values.
func (style Style) SetOuterSpacing(x, y Unit) {
	style.SetMarginLeft(css.Decode(x))
	style.SetMarginRight(css.Decode(x))

	style.SetMarginTop(css.Decode(y))
	style.SetMarginBottom(css.Decode(y))
}

//SetOuterSpacingTop sets the outer spacing top, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingTop(value Unit) {
	style.SetMarginTop(css.Decode(value))
}

//SetOuterSpacingLeft sets the outer spacing left, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingLeft(value Unit) {
	style.SetMarginLeft(css.Decode(value))
}

//SetOuterSpacingBottom sets the outer spacing bottom, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingBottom(value Unit) {
	style.SetMarginBottom(css.Decode(value))
}

//SetOuterSpacingRight sets the outer spacing right, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingRight(value Unit) {
	style.SetMarginRight(css.Decode(value))
}

//SetOffsetTop sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetTop(offset Unit) {
	style.SetTop(css.Decode(offset))
}

//SetOffsetBottom sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetBottom(offset Unit) {
	style.SetBottom(css.Decode(offset))
}

//SetOffsetLeft sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetLeft(offset Unit) {
	style.SetLeft(css.Decode(offset))
}

//SetOffsetRight sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetRight(offset Unit) {
	style.SetRight(css.Decode(offset))
}
