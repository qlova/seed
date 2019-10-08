package style

import "github.com/qlova/seed/style/css"

//TextAligner aligns text.
type TextAligner struct {
	style Style
}

//Left aligned text.
func (a TextAligner) Left() {
	a.style.SetTextAlign(css.Left)
}

//Right aligned text.
func (a TextAligner) Right() {
	a.style.SetTextAlign(css.Right)
}

//Center aligned text.
func (a TextAligner) Center() {
	a.style.SetTextAlign(css.Center)
}

//TextAlign returns a TextAligner for aligning text.
func (style Style) TextAlign() TextAligner {
	return TextAligner{style}
}

//Aligner aligns items.
type Aligner interface {
	Start()
	Center()
	End()
}

type selfAligner struct {
	Style
}

func (a selfAligner) Start() {
	a.SetAlignSelf(css.FlexStart)
}

func (a selfAligner) Center() {
	a.SetAlignSelf(css.Center)
}

func (a selfAligner) End() {
	a.SetAlignSelf(css.FlexEnd)
}

//Align returns an aligner that aligns the seed.
func (style Style) Align() Aligner {
	return selfAligner{style}
}

type itemsAligner struct {
	Style
}

func (a itemsAligner) Start() {
	a.SetAlignItems(css.FlexStart)
}

func (a itemsAligner) Center() {
	a.SetAlignItems(css.Center)
}

func (a itemsAligner) End() {
	a.SetAlignItems(css.FlexEnd)
}

//AlignItems returns an aligner that aligns the children of this seed.
func (style Style) AlignItems() Aligner {
	return selfAligner{style}
}

//SetLayer sets the rendering layer.
func (style Style) SetLayer(layer int) {
	style.SetZIndex(css.Integer(layer))
}

//Wrap causes the children elements of this element to wrap like text.
func (style Style) Wrap() {
	style.Style.SetFlexWrap(css.Wrap)
}
