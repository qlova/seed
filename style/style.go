package style

import (
	"github.com/qlova/seed/style/css"
)

//Style is a set of visual indications of an element.
//For example, colour, spacing & positioning.
type Style struct {
	Style css.Style

	x     *complex128
	y     *complex128
	angle *float64
	rx    *float64
	scale *float64

	selectors map[string]Style
}

//New returns a new Style.
func New() Style {
	return Style{
		Style: css.NewStyle(),
	}
}

//Select allows you to select custom style selectors such as :hover
func (s Style) Select(selector string) Style {
	if style, ok := s.selectors[selector]; ok {
		return style
	}
	var created = New()
	s.selectors[selector] = created
	return created
}

//From returns a style from a css stylable.
func From(stylable css.Stylable) Style {
	return Style{
		Style:     css.Style{stylable},
		selectors: make(map[string]Style, 0),
	}
}

//Copy duplicates a style and returns a copy of it.
func (style Style) Copy() Style {
	var OldStyleImplemenation = style.CSS().Stylable.(css.Implementation)
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
func (style Style) CSS() css.Style {
	return style.Style.CSS()
}

//Bytes return the style as CSS.
func (style Style) Bytes() []byte {

	style.update()

	return style.Style.Bytes()
}

//SetBold sets the text of this element to be bold.
func (style Style) SetBold() {
	style.CSS().SetFontWeight(css.Bold)
}

//SetHidden sets this to be hidden and removed from flow.
func (style Style) SetHidden() {
	style.CSS().SetDisplay(css.None)
}

//SetInvisible sets this to be invisible.
func (style Style) SetInvisible() {
	style.CSS().SetVisibility(css.Hidden)
}

//SetVisible sets this to be not hidden and visible.
func (style Style) SetVisible() {
	style.CSS().SetDisplay(css.Flex)
	style.CSS().SetVisibility(css.Visible)
}

//SetCol sets this element to behave like a column when rendering children (rendering them vertically).
func (style Style) SetCol() {
	style.CSS().SetFlexDirection(css.Column)
	style.CSS().SetDisplay(css.InlineFlex)
}

//SetRow sets this element to behave like a row when rendering children (rendering them horizontally).
func (style Style) SetRow() {
	style.CSS().SetFlexDirection(css.Row)
	style.CSS().SetDisplay(css.InlineFlex)
}
