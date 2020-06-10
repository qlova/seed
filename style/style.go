package style

import (
	"fmt"
	"image/color"

	"github.com/qlova/seed"

	"github.com/qlova/seed/css"
)

type Unit complex64

func (u Unit) String() string {
	if u == 0 {
		return "0"
	}
	return fmt.Sprintf("%v%%", real(u))
}

func (u Unit) Unit() css.Unit {
	if imag(u) == 0 {
		if real(u) == 0 {
			return css.Unit{}
		}
		return css.Percent(real(u))
	}

	var ratio = imag(u) / real(u)
	switch {
	case ratio > 0.9 && ratio < 1.1:
		return css.Em(real(u))
	case ratio > 1.9 && ratio < 2.1:
		return css.Vmin(real(u))
	case ratio > 2.9 && ratio < 3.1:
		return css.Rem(real(u))
	case ratio > 4.9 && ratio < 5.1:
		return css.Px(real(u))
	}

	return css.Unit{}
}

type Style interface {
	seed.Option
	Rules() css.Rules
}

//Translate the element by the given x and y values.
//This overrrides any previous calls to Translate.
func Translate(x, y Unit) css.Rules {
	return css.Rules{
		css.Set("--x", x.String()),
		css.Set("--y", y.String()),
		css.Set("transform", "translate(var(--x, 0), var(--y, 0)) rotate(var(--angle, 0)) scale(var(--scale, 1), var(--scale, 1))"),
	}
}

//Rotate the element by the given angle value.
//This overrrides any previous calls to Angle.
func Rotate(angle float32) css.Rules {
	return css.Rules{
		css.Set("--angle", fmt.Sprintf(`%frad`, angle)),
		css.Set("transform", "translate(var(--x, 0), var(--y, 0)) rotate(var(--angle, 0)) scale(var(--scale, 1), var(--scale, 1))"),
	}
}

//Scale the element by the given scale value.
//This overrrides any previous calls to Scale.
func Scale(scale float32) css.Rules {
	return css.Rules{
		css.Set("--scale", fmt.Sprintf(`%v`, scale)),
		css.Set("transform", "translate(var(--x, 0), var(--y, 0)) rotate(var(--angle, 0)) scale(var(--scale, 1), var(--scale, 1))"),
	}
}

//SetTextColor sets the color of the seed.
func SetTextColor(c color.Color) css.Rule {
	return css.SetColor(css.RGB{Color: c})
}

//SetColor sets the color of the seed.
func SetColor(c color.Color) css.Rule {
	return css.SetBackgroundColor(css.RGB{Color: c})
}

//SetColumn sets the seed to behave as a column.
func SetColumn() css.Rule {
	return css.SetFlexDirection(css.Column)
}

//SetRow sets the seed to behave as a column.
func SetRow() css.Rule {
	return css.SetFlexDirection(css.Row)
}

//SetHidden removes the seed.
func SetHidden() css.Rule {
	return css.SetDisplay(css.None)
}

//SetVisible sets the seed to be visible.
func SetVisible() css.Rule {
	return css.SetDisplay(css.Flex)
}

//SetHeight sets the height of the seed.
func SetHeight(height Unit) css.Rules {
	return css.Rules{
		css.SetHeight(height.Unit()),
	}
}

//SetWidth sets the width of the seed.
func SetWidth(width Unit) css.Rules {
	return css.Rules{
		css.SetWidth(width.Unit()),
	}
}

//SetSize sets the width and height of the seed.
func SetSize(width, height Unit) css.Rules {
	return append(
		SetWidth(width),
		SetHeight(height)...,
	)
}

//SetMaxHeight sets the height of the seed.
func SetMaxHeight(height Unit) css.Rule {
	return css.SetMaxHeight(height.Unit())
}

//SetMaxWidth sets the width of the seed.
func SetMaxWidth(width Unit) css.Rule {
	return css.SetMaxWidth(width.Unit())
}

//SetMinHeight sets the minumum height of the seed.
func SetMinHeight(height Unit) css.Rule {
	return css.SetMinHeight(height.Unit())
}

//SetMinWidth sets the minumum width of the seed.
func SetMinWidth(width Unit) css.Rule {
	return css.SetMinWidth(width.Unit())
}

//SetLayer sets the layer of the seed.
func SetLayer(layer int) css.Rule {
	return css.SetZIndex(css.Int(layer))
}

//SetOpacity sets the seed's opacity.
func SetOpacity(opacity float32) css.Rule {
	return css.SetOpacity(css.Number(opacity))
}

//SetRoundedCorners sets the seed to have rounded corners of the specified sizing.
func SetRoundedCorners(first Unit, size ...Unit) css.Rule {
	if len(size) == 0 {
		return css.SetBorderRadius(first.Unit())
	}
	if len(size) == 3 {
		return css.Set("border-radius", fmt.Sprintf("%v %v %v %v", first.Unit().Rule(), size[0].Unit().Rule(), size[1].Unit().Rule(), size[2].Unit().Rule()))
	}
	panic("invalid arguments")
}

//Expand sets the seed to expand to fill up space.
func Expand() css.Rule {
	return css.SetFlexGrow(css.Number(1))
}

//Shrink sets the seed to shrink if needed to save space.
func Shrink() css.Rule {
	return css.SetFlexShrink(css.Number(1))
}

//Wrap sets the contents to wrap.
func Wrap() css.Rule {
	return css.SetFlexWrap(css.Wrap)
}

//SetScrollable sets that this can be scrolled vertically.
func SetScrollable() css.Rules {
	return css.Rules{
		css.SetOverflowY(css.Auto),
		css.SetOverflowX(css.Hidden),
		css.Set("-webkit-overflow-scrolling", "touch"),
		css.Set("-webkit-overscroll-behavior", "contain"),
		css.Set("overscroll-behavior", "contain"),
	}
}

//Clip sets that contents cannot overflow.
func Clip() css.Rules {
	return css.Rules{
		css.SetOverflow(css.Hidden),
	}
}
