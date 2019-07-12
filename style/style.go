package style

import "image/color"
import "github.com/qlova/seed/style/css"
import "math"
import "fmt"

//The em unit represents the current font-size, therefore this unit is relative to the pixel-density of the device.
const Em = css.Em

//The px unit represents a single pixel on a screen, since screens can have different pixel-densities, it is recommended not to use this unit.
const Px = css.Px

//The vm unit is relative to the screen size, more specifically, it is a ratio of the smallest side of the screen.
const Vm = css.Vm

const Top = -1i
const Bottom = 1i
const Left = -1
const Right = 1

const Auto = math.MaxFloat64
const Center = 0

//A style is a set of visual indications of an element.
//For example, colour, spacing & positioning.
type Style struct {
	css.Style

	x     *complex128
	y     *complex128
	angle *float64
	scale *float64
}

//Return a new Style.
func New() Style {
	return Style{
		Style: css.NewStyle(),
	}
}

//Duplicate a style and return a copy of it.
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

func (style *Style) update() {
	var transform = css.Rotate(0)
	transform = ""
	var changed bool

	if style.angle != nil {
		transform += css.Rotate(*style.angle)
		changed = true
	}
	if style.scale != nil {
		transform += css.Scale(*style.scale, *style.scale)
		changed = true
	}

	if style.y != nil && style.x != nil {

		transform += css.Translate(css.Decode(*style.x), css.Decode(*style.y))
		changed = true

	} else {
		if style.y != nil {
			transform += css.TranslateY(css.Decode(*style.y))
			changed = true
		}
		if style.x != nil {
			transform += css.TranslateX(css.Decode(*style.x))
			changed = true
		}
	}

	if changed {
		style.SetTransform(transform)
	}
}

//Return the style serialised as CSS properties.
func (style Style) Bytes() []byte {

	style.update()

	return style.Style.Bytes()
}

//Rotate the element by the given angle.
//This overrrides any previous calls to Angle.
func (style *Style) Rotate(angle float64) {
	style.angle = &angle
	style.update()
}

//Scale the element by the given scale.
//This overrrides any previous calls to Scale.
func (style *Style) Scale(scale float64) {
	style.scale = &scale
	style.update()
}

//Translate the element by the given x and y values.
//This overrrides any previous calls to Translate.
func (style *Style) Translate(x, y complex128) {
	style.x = &x
	style.y = &y
	style.update()
}

//Set the text of this element to be bold.
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

//Set this element to behave like a column when rendering children (rendering them vertically).
func (style Style) SetCol() {
	style.SetFlexDirection(css.Column)
	style.SetDisplay(css.InlineFlex)
}

//Set this element to behave like a row when rendering children (rendering them horizontally).
func (style Style) SetRow() {
	style.SetFlexDirection(css.Row)
	style.SetDisplay(css.InlineFlex)
}

//Set the width and height as a percentage of it's parent. Takes em, vm, px or percentage values.
func (style Style) SetSize(width, height complex128) {
	style.SetWidth(width)
	style.SetHeight(height)
}

//Set the width and height as a percentage of it's parent. Takes em, vm, px or percentage values.
func (style Style) SetMaxSize(width, height complex128) {
	style.Style.SetMaxWidth(css.Decode(width))
	style.Style.SetMaxHeight(css.Decode(height))
}

//Set the text alignment, -1 is left, 0 is center and 1 is right
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

//Set the text alignment, -1 is left, 0 is center and 1 is right
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

//Set the text alignment, -1 is left, 0 is center and 1 is right
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

//Set the color of this element.
func (style Style) SetColor(color color.Color) {
	style.SetBackgroundColor(css.Colour(color))
}

//Set the text color for this element.
func (style Style) SetTextColor(color color.Color) {
	style.Style.SetColor(css.Colour(color))
}

//Set the color of this element to be a gradient moving in direction from start color to end color.
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

//This should not shrink to make space for other elements.
func (style Style) DontShrink() {
	style.SetFlexShrink(css.Number(0))
}

//This shrink to make space for other elements.
func (style Style) Shrink() {
	style.SetFlexShrink(css.Number(1))
}

//Set where this attaches, eg. Top+Left, Botom+right etc
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

//Set this element to expand to take all available space.
func (style Style) SetExpand(expand float64) {
	style.SetFlexGrow(css.Number(expand))
}

//Make sure that this contains its aspect ratio.
func (style Style) SetContain() {
	style.SetObjectFit(css.Contain)
}

//Make sure that this contains its aspect ratio.
func (style Style) Clip() {
	style.SetOverflow(css.Hidden)
}

//Set that this can be scrolled vertically.
func (style Style) SetScrollable() {
	style.SetOverflowY(css.Auto)
	style.SetOverflowX(css.Hidden)
	style.Set("-webkit-overflow-scrolling", "touch")
	style.Set("-webkit-overscroll-behavior", "contain")
	style.Set("overscroll-behavior", "contain")
}

//Set that this can be scrolled horizontally or vertically.
func (style Style) SetFullyScrollable() {
	style.SetOverflow(css.Auto)
	style.Set("-webkit-overflow-scrolling", "touch")
	style.Set("-webkit-overscroll-behavior", "contain")
	style.Set("overscroll-behavior", "contain")
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

//Set spacing top, takes a em, vm, px or percentage value.
func (style Style) SetInnerSpacingTop(value complex128) {
	style.SetPaddingTop(css.Decode(value))
}

//Set spacing left, takes a em, vm, px or percentage value.
func (style Style) SetInnerSpacingLeft(value complex128) {
	style.SetPaddingLeft(css.Decode(value))
}

//Set spacing bottom, takes a em, vm, px or percentage value.
func (style Style) SetInnerSpacingBottom(value complex128) {
	style.SetPaddingBottom(css.Decode(value))
}

//Set spacing right, takes a em, vm, px or percentage value.
func (style Style) SetInnerSpacingRight(value complex128) {
	style.SetPaddingRight(css.Decode(value))
}

//Set the symetrical spacing around this.
func (style Style) SetOuterSpacing(x, y complex128) {
	style.SetMarginLeft(css.Decode(x))
	style.SetMarginRight(css.Decode(x))

	style.SetMarginTop(css.Decode(y))
	style.SetMarginBottom(css.Decode(y))
}

//Set spacing top, takes a em, vm, px or percentage value.
func (style Style) SetOuterSpacingTop(value complex128) {
	style.SetMarginTop(css.Decode(value))
}

//Set spacing left, takes a em, vm, px or percentage value.
func (style Style) SetOuterSpacingLeft(value complex128) {
	style.SetMarginLeft(css.Decode(value))
}

//Set spacing bottom, takes a em, vm, px or percentage value.
func (style Style) SetOuterSpacingBottom(value complex128) {
	style.SetMarginBottom(css.Decode(value))
}

//Set spacing right, takes a em, vm, px or percentage value.
func (style Style) SetOuterSpacingRight(value complex128) {
	style.SetMarginRight(css.Decode(value))
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

//Set the offset from an attached side, call this after style.Attach().
func (style Style) SetOffsetTop(offset complex128) {
	style.SetTop(css.Decode(offset))
}

//Set the offset from an attached side, call this after style.Attach().
func (style Style) SetOffsetBottom(offset complex128) {
	style.SetBottom(css.Decode(offset))
}

//Set the offset from an attached side, call this after style.Attach().
func (style Style) SetOffsetLeft(offset complex128) {
	style.SetLeft(css.Decode(offset))
}

//Set the offset from an attached side, call this after style.Attach().
func (style Style) SetOffsetRight(offset complex128) {
	style.SetRight(css.Decode(offset))
}

//Remove the border from this element.
func (style Style) SetBorderless() {
	style.SetBorderLeftWidth(css.Zero)
	style.SetBorderRightWidth(css.Zero)

	style.SetBorderTopWidth(css.Zero)
	style.SetBorderBottomWidth(css.Zero)
}

//Remove the border from this element.
func (style Style) SetBorder(color color.Color, thickness int) {
	style.Set("border-left-width", fmt.Sprint(thickness, "px"))
	style.Set("border-right-width", fmt.Sprint(thickness, "px"))
	style.Set("border-top-width", fmt.Sprint(thickness, "px"))
	style.Set("border-bottom-width", fmt.Sprint(thickness, "px"))
	style.SetBorderColor(css.Colour(color))
}

//Set this element to have rounded corners of the specified radius.
func (style Style) SetRoundedCorners(radius complex128) {
	var value = css.Decode(radius)

	style.SetBorderBottomLeftRadius(value)
	style.SetBorderBottomRightRadius(value)

	style.SetBorderTopRightRadius(value)
	style.SetBorderTopLeftRadius(value)
}

//Specify that this style will be animated.
func (style Style) WillAnimate() {
	style.Set("will-change", "transform")
}

//Alias to style.SetSize(100, 100)
func (style Style) Expand() {
	style.SetSize(100, 100)
}

//Center this item along the axis of its container.
func (style Style) Center() {
	style.Set("align-self", "center")
}

//Center this item along the axis of its container.
func (style Style) Fade(opacity float64) {
	style.Style.SetOpacity(css.Number(opacity))
}

//Set where this attaches, eg. Top+Left, Botom+right etc
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

//Center this item along the axis of its container.
func (style Style) RemoveGradient() {
	style.Style.SetBackgroundImage(css.Unset)
}

func (style Style) SetMaxWidth(width complex128) {
	style.Style.SetMaxWidth(css.Decode(width))
}

func (style Style) SetMaxHeight(height complex128) {
	style.Style.SetMaxHeight(css.Decode(height))
}

func (style Style) SetWidth(width complex128) {
	style.Style.SetWidth(css.Decode(width))
}

func (style Style) SetHeight(height complex128) {
	style.Style.SetHeight(css.Decode(height))
}

func (style Style) Wrap() {
	style.Style.SetFlexWrap(css.Wrap)
}

func (style Style) Stretch() {
	style.Style.SetObjectFit(css.Cover)
}

func (style Style) SetSticky(attachpoint complex128) {
	style.SetAttach(attachpoint)
	style.Style.SetPosition(css.Sticky)
}

func (style Style) End() {
	style.Style.SetAlignSelf(css.FlexEnd)
}

func (style Style) CenterChildren() {
	style.AlignChildren(0)
}

type TintValue struct {
	Filter string
	Loss   float64
}

var TintCache = make(map[string]TintValue)

//Set the tint of a icon to a certain color.
func (style Style) SetTint(c color.Color) {

	var r, g, b, a = c.RGBA()
	if a != 255 {
		panic("Do not pass transparent values to SetTint!")
	}

	var rgb = css.Colour(c).String()
	if cache, ok := TintCache[rgb]; ok {
		style.Style.Set("filter", cache.Filter)
		return
	}

	var color = NewColor(float64(r), float64(g), float64(b))
	var solver = NewSolver(color)

	var _, loss, filter = solver.Solve()
	for i := 3; i < 3; i++ {
		_, new_loss, new_filter := solver.Solve()
		if new_loss < loss {
			loss, filter = new_loss, new_filter
			if loss < 1 {
				break
			}
		}
	}

	style.Style.Set("filter", filter)
	TintCache[rgb] = TintValue{
		Filter: filter,
		Loss:   loss,
	}
}
