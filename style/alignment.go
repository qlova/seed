package style

import "github.com/qlova/seed/style/css"

//SetLayer sets the rendering layer.
func (style Style) SetLayer(layer int) {
	style.SetZIndex(css.Integer(layer))
}

//TODO fix these methods, they are a complete mess and unclear.

//TextAlign sets the text alignment, -1 is left, 0 is center and 1 is right
func (style Style) TextAlign(alignment float64) {
	switch alignment {
	case 0:
		style.SetTextAlign(css.Center)
	case -1:
		style.SetTextAlign(css.Left)
	case 1:
		style.SetTextAlign(css.Right)
	}
}

//Align sets the alignment, -1 is left, 0 is center and 1 is right
func (style Style) Align(alignment float64) {
	switch alignment {
	case 0:
		style.SetTextAlign(css.Center)
		style.SetAlignSelf(css.Center)
	case -1:
		style.SetTextAlign(css.Left)
		style.SetAlignSelf(css.FlexStart)
	case 1:
		style.SetTextAlign(css.Right)
		style.SetAlignSelf(css.FlexEnd)
	}
}

//AlignChildren aligns the children of this element, -1 is left, 0 is center and 1 is right
func (style Style) AlignChildren(alignment float64) {
	switch alignment {
	case 0:
		style.SetJustifyContent(css.Center)
	case -1:
		style.SetJustifyContent(css.FlexStart)
	case 1:
		style.SetJustifyContent(css.FlexEnd)
	}
}

//SetAlignment sets the text alignment, -1 is left, 0 is center and 1 is right
func (style Style) SetAlignment(align float64) {
	switch align {
	case 0:
		style.SetTextAlign(css.Center)
		style.SetAlignSelf(css.Center)
		//style.SetJustifySelf(css.Center)
		style.Set("justify-self", "center")
	case -1:
		style.SetTextAlign(css.Left)
		style.SetAlignSelf(css.FlexStart)
		//style.SetJustifySelf(css.FlexStart)
		style.Set("justify-self", "flex-start")
	case 1:
		style.SetTextAlign(css.Right)
		style.SetAlignSelf(css.FlexEnd)
		//style.SetJustifySelf(css.FlexEnd)
		style.Set("justify-self", "flex-end")
	}
}

//SetChildAlignment sets the child alignment, -1 is left, 0 is center and 1 is right
func (style Style) SetChildAlignment(align float64) {
	switch align {
	case 0:
		style.SetTextAlign(css.Center)
		style.SetAlignContent(css.Center)
		style.SetAlignItems(css.Center)
		style.SetJustifyContent(css.Center)
	case -1:
		style.SetTextAlign(css.Left)
		style.SetAlignContent(css.FlexStart)
		style.SetJustifyContent(css.FlexStart)
	case 1:
		style.SetTextAlign(css.Right)
		style.SetAlignContent(css.FlexEnd)
		style.SetJustifyContent(css.FlexEnd)
	}
}

//SetAttach sets where this attaches, eg. Top+Left, Botom+right etc
func (style Style) SetAttach(attach complex128) {
	switch real(attach) {
	case -1:
		style.SetLeft(css.Zero)
		style.SetPosition(css.Fixed)
	case 0:
		style.SetLeft(css.Initial)
		style.SetRight(css.Initial)
	case 1:
		style.SetRight(css.Zero)
		style.SetPosition(css.Fixed)
	}
	switch imag(attach) {
	case -1:
		style.SetTop(css.Zero)
		style.SetPosition(css.Fixed)
	case 0:
		style.SetTop(css.Initial)
		style.SetBottom(css.Initial)
	case 1:
		style.SetBottom(css.Zero)
		style.SetPosition(css.Fixed)
	}
}

//AttachToParent sets where this attaches, eg. Top+Left, Botom+right etc
//For example, if the attachpoint is right, then this object's right side is attached to its parent's right side.
func (style Style) AttachToParent(attachpoint complex128) {
	switch real(attachpoint) {
	case -1:
		style.SetLeft(css.Zero)
		style.SetPosition(css.Absolute)
	case 0:
		style.SetLeft(css.Initial)
		style.SetRight(css.Initial)
	case 1:
		style.SetRight(css.Zero)
		style.SetPosition(css.Absolute)
	}
	switch imag(attachpoint) {
	case -1:
		style.SetTop(css.Zero)
		style.SetPosition(css.Absolute)
	case 0:
		style.SetTop(css.Initial)
		style.SetBottom(css.Initial)
	case 1:
		style.SetBottom(css.Zero)
		style.SetPosition(css.Absolute)
	}
}

//SetSticky is like SetAttach but the element will be sticky to the screen when scrolling.
func (style Style) SetSticky(attachpoint complex128) {
	style.SetAttach(attachpoint)
	style.Style.SetPosition(css.Sticky)
}

//Wrap causes the children elements of this element to wrap like text.
func (style Style) Wrap() {
	style.Style.SetFlexWrap(css.Wrap)
}

//Start aligns this element to the start of its container.
func (style Style) Start() {
	style.Style.SetAlignSelf(css.FlexStart)
}

//End aligns this element to the end of its container.
func (style Style) End() {
	style.Style.SetAlignSelf(css.FlexEnd)
}

//CenterChildren centers the children of this element.
func (style Style) CenterChildren() {
	style.AlignChildren(0)
}

//Center this item along the axis of its container.
func (style Style) Center() {
	style.Set("align-self", "center")
}
