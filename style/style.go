package style

import "image/color"
import "github.com/qlova/seed/style/css"
import "math"
import "math/big"
import "encoding/base64"

const Em = css.Em
const Px = css.Px

const Top = -1i
const Bottom = 1i
const Left = -1
const Right = 1

const Auto = 0
const Center = 0

type Style struct {
	css.Style

	angle *float64
	scale *float64
}

func New() Style {
	return Style{
		Style: css.NewStyle(),
	}
}

type Font struct {
	name string
	css.FontFace
}

var font_id int64 = 1;
func NewFont(path string) Font {
	
	var id = base64.RawURLEncoding.EncodeToString(big.NewInt(font_id).Bytes())
	font_id++
	
	var font = Font{
		name: id,
		FontFace: css.NewFontFace(id, path),
	}

	//Avoid invisisible text while webfonts are loading.
	//font.FontFace.FontDisplay = css.Swap

	return font
}

func (style Style) Bytes() []byte {
	var transform = css.Rotate(0)
	var changed bool
	
	if style.angle != nil {
		transform += css.Rotate(*style.angle)
		changed = true
	}
	if style.scale != nil {
		transform += css.Scale(*style.scale, *style.scale)
		changed = true
	}

	if changed {
		style.SetTransform(transform)
	}

	return style.Style.Bytes()
}

func (style *Style) Rotate(angle float64) {
	style.angle = &angle
}
func (style *Style) Scale(scale float64) {
	style.scale = &scale
}

//Set the symetrical spacing within this.
func (style Style) SetFont(font Font) {
	style.SetFontFamily(font.FontFace)
}

//Set the symetrical spacing within this.
func (style Style) SetBold() {
	style.SetFontWeight(css.Bold)
}

//Set the Text Size, a multiple of the default text size.
func (style Style) SetTextSize(size complex128) {
	style.SetFontSize(css.Decode(size))
}

//Set this to be hidden.
func (style Style) SetHidden() {
	style.SetDisplay(css.None)
}

//Set this to be visible.
func (style Style) SetVisible() {
	style.SetDisplay(css.Flex)
}

//Set the width and height as a percentage of it's parent. A value of 0 means it is calculated automatically.
func (style Style) SetSize(width, height complex128) {
	style.SetWidth(css.Decode(width))
	style.SetHeight(css.Decode(height))
}

//Set the width and height as a percentage of it's parent. A value of 0 means it is calculated automatically.
func (style Style) SetMaxSize(width, height complex128) {
	style.SetMaxWidth(css.Decode(width))
	style.SetMaxHeight(css.Decode(height))
}

//Set the text alignment, -1 is left, 0 is center and 1 is right
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

//Set the text alignment, -1 is left, 0 is center and 1 is right
func (style Style) SetChildAlignment(align float64) {
	switch align {
		case 0:
			style.SetTextAlign(css.Center)
			style.SetAlignContent(css.Center)
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

//Set the text alignment, -1 is left, 0 is center and 1 is right
func (style Style) SetColor(color color.Color) {
	style.SetBackgroundColor(css.Colour(color))
}

//Set the text alignment, -1 is left, 0 is center and 1 is right
func (style Style) SetTextColor(color color.Color) {
	style.Style.SetColor(css.Colour(color))
}

//Set the text alignment, -1 is left, 0 is center and 1 is right
func (style Style) SetGradient(direction complex128, start, end color.Color) {
	style.SetBackgroundImage(css.LinearGradient(math.Atan2(imag(direction), real(direction))+math.Pi/2, css.Colour(start), css.Colour(end)))
}

//Set the rendering layer, this is the order that this will be rendered in.
func (style Style) SetLayer(layer int) {
	style.SetZIndex(css.Integer(layer))
}

//This should not shrink to make space for other elements.
func (style Style) SetUnshrinkable() {
	style.SetFlexShrink(css.Number(0))
}

//Set where this attaches to, 0 0 is unattached, -1, 0 is attached to left, 1 1 is attached to bottom right etc.
func (style Style) SetAttach(attach complex64) {
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

//Set the rendering layer, this is the order that this will be rendered in.
func (style Style) SetExpand(expand float64) {
	style.SetFlexGrow(css.Number(expand))
}

//Make sure that this contains its aspect ratio.
func (style Style) SetContain() {
	style.SetObjectFit(css.Contain)
}

//Set that this can be scrolled.
func (style Style) SetScrollable() {
	style.SetOverflow(css.Auto)
}

//Set that this cannot be scrolled.
func (style Style) SetNotScrollable() {
	style.SetOverflow(css.Hidden)
}

//Set the symetrical spacing within this.
func (style Style) SetInnerSpacing(x, y complex128) {
	style.SetPaddingLeft(css.Decode(x))
	style.SetPaddingRight(css.Decode(x))
	
	style.SetPaddingTop(css.Decode(y))
	style.SetPaddingBottom(css.Decode(y))
}

//Set the symetrical spacing within this.
func (style Style) SetOuterSpacing(x, y complex128) {
	style.SetMarginLeft(css.Decode(x))
	style.SetMarginRight(css.Decode(x))
	
	style.SetMarginTop(css.Decode(y))
	style.SetMarginBottom(css.Decode(y))
}

//Set the offset from an attached side, call this after style.Attach().
func (style Style) SetOffset(side complex128, offset complex128) {
	switch side {
		case Left:
			style.SetLeft(css.Decode(offset))
		case Right:
			style.SetRight(css.Decode(offset))
		case Top:
			style.SetTop(css.Decode(offset))
		case Bottom:
			style.SetBottom(css.Decode(offset))
	}
}

//Set the symetrical spacing within this.
func (style Style) SetBorderless() {
	style.SetBorderLeftWidth(css.Zero)
	style.SetBorderRightWidth(css.Zero)
	
	style.SetBorderTopWidth(css.Zero)
	style.SetBorderBottomWidth(css.Zero)
}

//Set the symetrical spacing within this.
func (style Style) SetRoundedCorners(radius complex128) {
	var value = css.Decode(radius)
	
	style.SetBorderBottomLeftRadius(value)
	style.SetBorderBottomRightRadius(value)
	
	style.SetBorderTopRightRadius(value)
	style.SetBorderTopLeftRadius(value)
}
