package style

import (
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
)

//SetInnerSpacing sets the spacing within this element, takes em, vm, px or percentage values.
func (style Style) SetInnerSpacing(x, y  unit.Unit) {
	style.SetPaddingLeft(css.Decode(x))
	style.SetPaddingRight(css.Decode(x))

	style.SetPaddingTop(css.Decode(y))
	style.SetPaddingBottom(css.Decode(y))
}

//SetInnerSpacingTop sets the inner spacing top, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingTop(value  unit.Unit) {
	style.SetPaddingTop(css.Decode(value))
}

//SetInnerSpacingLeft sets the inner spacing left, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingLeft(value  unit.Unit) {
	style.SetPaddingLeft(css.Decode(value))
}

//SetInnerSpacingBottom sets the inner spacing bottom, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingBottom(value  unit.Unit) {
	style.SetPaddingBottom(css.Decode(value))
}

//SetInnerSpacingRight sets the inner spacing right, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingRight(value  unit.Unit) {
	style.SetPaddingRight(css.Decode(value))
}

//SetOuterSpacing sets the outer spacing of this element, takes em, vm, px or percentage values.
func (style Style) SetOuterSpacing(x, y  unit.Unit) {
	style.SetMarginLeft(css.Decode(x))
	style.SetMarginRight(css.Decode(x))

	style.SetMarginTop(css.Decode(y))
	style.SetMarginBottom(css.Decode(y))
}

//SetOuterSpacingTop sets the outer spacing top, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingTop(value  unit.Unit) {
	style.SetMarginTop(css.Decode(value))
}

//SetOuterSpacingLeft sets the outer spacing left, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingLeft(value  unit.Unit) {
	style.SetMarginLeft(css.Decode(value))
}

//SetOuterSpacingBottom sets the outer spacing bottom, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingBottom(value  unit.Unit) {
	style.SetMarginBottom(css.Decode(value))
}

//SetOuterSpacingRight sets the outer spacing right, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingRight(value  unit.Unit) {
	style.SetMarginRight(css.Decode(value))
}

//SetOffsetTop sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetTop(offset  unit.Unit) {
	style.SetTop(css.Decode(offset))
}

//SetOffsetBottom sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetBottom(offset  unit.Unit) {
	style.SetBottom(css.Decode(offset))
}

//SetOffsetLeft sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetLeft(offset  unit.Unit) {
	style.SetLeft(css.Decode(offset))
}

//SetOffsetRight sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetRight(offset  unit.Unit) {
	style.SetRight(css.Decode(offset))
}

//Spacer distributes spacing.
type Spacer interface {
	//Before provides spacing before children. This is like placing an expander before the children.
	Before()

	//Outside provides spacing outside of the children, effectively centering the children.
	//This is like placing an expander before and after the children.
	Outside()

	//Center provides spacing that centers the children.
	Center()

	//After provides spacing after children. This is like placing an expander after the children.
	After()

	//Inside distributes spacing inside children. This is like placing an expander between children.
	Inside()

	//Around distributes spacing around the children. This is like placing an expander before and after every child.
	Around()

	//Divide distributes the spacing evenly. This is like placing an expander before the children, between each child and after the children.
	Divide()
}

type itemSpacer struct {
	Style
}

func (s itemSpacer) Before() {
	s.SetJustifyContent(css.FlexStart)
}

func (s itemSpacer) Outside() {
	s.SetJustifyContent(css.Center)
}

func (s itemSpacer) Center() {
	s.Outside()
}

func (s itemSpacer) After() {
	s.SetJustifyContent(css.FlexEnd)
}

func (s itemSpacer) Inside() {
	s.SetJustifyContent(css.SpaceBetween)
}

func (s itemSpacer) Around() {
	s.SetJustifyContent(css.SpaceAround)
}

func (s itemSpacer) Divide() {
	s.SetJustifyContent(css.SpaceEvenly)
}

//ItemSpacing returns a spacer for distributing spacing to children.
func (style Style) ItemSpacing() Spacer {
	return itemSpacer{style}
}

type wrapSpacer struct {
	Style
}

func (s wrapSpacer) Before() {
	s.SetAlignContent(css.FlexStart)
}

func (s wrapSpacer) Outside() {
	s.SetAlignContent(css.Center)
}

func (s wrapSpacer) Center() {
	s.Outside()
}

func (s wrapSpacer) After() {
	s.SetAlignContent(css.FlexEnd)
}

func (s wrapSpacer) Inside() {
	s.SetAlignContent(css.SpaceBetween)
}

func (s wrapSpacer) Around() {
	s.SetAlignContent(css.SpaceAround)
}

func (s wrapSpacer) Divide() {
	s.SetAlignContent(css.SpaceEvenly)
}

//WrapSpacing distributes spacing to wrapped lines of children of this seed..
func (style Style) WrapSpacing() Spacer {
	return wrapSpacer{style}
}
