package style

import (
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
)

//SetInnerSpacing sets the spacing within this element, takes em, vm, px or percentage values.
func (style Style) SetInnerSpacing(x, y unit.Unit) {
	style.CSS().SetPaddingLeft(css.Decode(x))
	style.CSS().SetPaddingRight(css.Decode(x))

	style.CSS().SetPaddingTop(css.Decode(y))
	style.CSS().SetPaddingBottom(css.Decode(y))
}

//SetInnerSpacingTop sets the inner spacing top, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingTop(value unit.Unit) {
	style.CSS().SetPaddingTop(css.Decode(value))
}

//SetInnerSpacingLeft sets the inner spacing left, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingLeft(value unit.Unit) {
	style.CSS().SetPaddingLeft(css.Decode(value))
}

//SetInnerSpacingBottom sets the inner spacing bottom, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingBottom(value unit.Unit) {
	style.CSS().SetPaddingBottom(css.Decode(value))
}

//SetInnerSpacingRight sets the inner spacing right, takes a em, vm, px or a percentage value.
func (style Style) SetInnerSpacingRight(value unit.Unit) {
	style.CSS().SetPaddingRight(css.Decode(value))
}

//SetOuterSpacing sets the outer spacing of this element, takes em, vm, px or percentage values.
func (style Style) SetOuterSpacing(x, y unit.Unit) {
	style.CSS().SetMarginLeft(css.Decode(x))
	style.CSS().SetMarginRight(css.Decode(x))

	style.CSS().SetMarginTop(css.Decode(y))
	style.CSS().SetMarginBottom(css.Decode(y))
}

//SetOuterSpacingTop sets the outer spacing top, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingTop(value unit.Unit) {
	style.CSS().SetMarginTop(css.Decode(value))
}

//SetOuterSpacingLeft sets the outer spacing left, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingLeft(value unit.Unit) {
	style.CSS().SetMarginLeft(css.Decode(value))
}

//SetOuterSpacingBottom sets the outer spacing bottom, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingBottom(value unit.Unit) {
	style.CSS().SetMarginBottom(css.Decode(value))
}

//SetOuterSpacingRight sets the outer spacing right, takes a em, vm, px or a percentage value.
func (style Style) SetOuterSpacingRight(value unit.Unit) {
	style.CSS().SetMarginRight(css.Decode(value))
}

//SetOffsetTop sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetTop(offset unit.Unit) {
	style.CSS().SetTop(css.Decode(offset))
}

//SetOffsetBottom sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetBottom(offset unit.Unit) {
	style.CSS().SetBottom(css.Decode(offset))
}

//SetOffsetLeft sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetLeft(offset unit.Unit) {
	style.CSS().SetLeft(css.Decode(offset))
}

//SetOffsetRight sets the offset from an attached side, call this after an Attach().
func (style Style) SetOffsetRight(offset unit.Unit) {
	style.CSS().SetRight(css.Decode(offset))
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
	s.CSS().SetJustifyContent(css.FlexStart)
}

func (s itemSpacer) Outside() {
	s.CSS().SetJustifyContent(css.Center)
}

func (s itemSpacer) Center() {
	s.Outside()
}

func (s itemSpacer) After() {
	s.CSS().SetJustifyContent(css.FlexEnd)
}

func (s itemSpacer) Inside() {
	s.CSS().SetJustifyContent(css.SpaceBetween)
}

func (s itemSpacer) Around() {
	s.CSS().SetJustifyContent(css.SpaceAround)
}

func (s itemSpacer) Divide() {
	s.CSS().SetJustifyContent(css.SpaceEvenly)
}

//ItemSpacing returns a spacer for distributing spacing to children.
func (style Style) ItemSpacing() Spacer {
	return itemSpacer{style}
}

type wrapSpacer struct {
	Style
}

func (s wrapSpacer) Before() {
	s.CSS().SetAlignContent(css.FlexStart)
}

func (s wrapSpacer) Outside() {
	s.CSS().SetAlignContent(css.Center)
}

func (s wrapSpacer) Center() {
	s.Outside()
}

func (s wrapSpacer) After() {
	s.CSS().SetAlignContent(css.FlexEnd)
}

func (s wrapSpacer) Inside() {
	s.CSS().SetAlignContent(css.SpaceBetween)
}

func (s wrapSpacer) Around() {
	s.CSS().SetAlignContent(css.SpaceAround)
}

func (s wrapSpacer) Divide() {
	s.CSS().SetAlignContent(css.SpaceEvenly)
}

//WrapSpacing distributes spacing to wrapped lines of children of this seed..
func (style Style) WrapSpacing() Spacer {
	return wrapSpacer{style}
}
