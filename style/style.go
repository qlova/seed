package style

import (
	"github.com/qlova/seed/style/css"
)

//Style is a set of visual indications of an element.
//For example, colour, spacing & positioning.
type Style struct {
	css.Style

	x     *complex128
	y     *complex128
	angle *float64
	rx    *float64
	scale *float64
}

//New returns a new Style.
func New() Style {
	return Style{
		Style: css.NewStyle(),
	}
}

//From returns a style from a css stylable.
func From(stylable css.Stylable) Style {
	return Style{
		Style: css.Style{stylable},
	}
}

//Copy duplicates a style and returns a copy of it.
func (style Style) Copy() Style {
	var OldStyleImplemenation = style.Stylable.(css.Implementation)
	var NewStyleImplementation = make(css.Implementation, len(OldStyleImplemenation))

	for key := range OldStyleImplemenation {
		NewStyleImplementation[key] = OldStyleImplemenation[key]
	}
	return Style{
		Style: css.Style{
			Stylable: NewStyleImplementation,
		},
	}
}

//Bytes return the style as CSS.
func (style Style) Bytes() []byte {

	style.update()

	return style.Style.Bytes()
}

//SetBold sets the text of this element to be bold.
func (style Style) SetBold() {
	style.SetFontWeight(css.Bold)
}

//SetHidden sets this to be hidden and removed from flow.
func (style Style) SetHidden() {
	style.SetDisplay(css.None)
}

//SetInvisible sets this to be invisible.
func (style Style) SetInvisible() {
	style.SetVisibility(css.Hidden)
}

//SetVisible sets this to be not hidden and visible.
func (style Style) SetVisible() {
	style.SetDisplay(css.Flex)
	style.SetVisibility(css.Visible)
}

//SetCol sets this element to behave like a column when rendering children (rendering them vertically).
func (style Style) SetCol() {
	style.SetFlexDirection(css.Column)
	style.SetDisplay(css.InlineFlex)
}

//SetRow sets this element to behave like a row when rendering children (rendering them horizontally).
func (style Style) SetRow() {
	style.SetFlexDirection(css.Row)
	style.SetDisplay(css.InlineFlex)
}
