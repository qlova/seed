package css

import (
	"fmt"
	color_go "image/color"
	"time"

	"qlova.org/seed/units"
)

type Duration time.Duration

func (Duration) durationValue() {}
func (d Duration) Rule() Rule {
	return Rule(fmt.Sprintf("%fs", time.Duration(d).Seconds()))
}

type AnimationName string

func (AnimationName) animationNameValue() {}
func (a AnimationName) Rule() Rule {
	return Rule(a)
}

type Number float64

func (Number) numberValue() {}
func (n Number) Rule() Rule {
	return Rule(fmt.Sprintf("%f", n))
}

type Int int

func (Int) integerOrAutoValue()           {}
func (Int) animationIterationCountValue() {}
func (i Int) Rule() Rule {
	return Rule(fmt.Sprintf("%v", i))
}

//Colour is the CSS color type.
type RGB struct {
	color_go.Color
}

func (RGB) colorValue() {}
func (c RGB) Rule() Rule {
	var col = color_go.NRGBAModel.Convert(c).(color_go.NRGBA)
	var r, g, b, a = col.R, col.G, col.B, col.A
	if a != 255 {
		return Rule(fmt.Sprint("rgba(", r, ",", g, ",", b, ",", float64(a)/255, ")"))
	} else {
		return Rule(fmt.Sprint("rgb(", r, ",", g, ",", b, ")"))
	}
}

type Unit struct {
	string
}

func Measure(u units.Unit) Unit {
	if u == nil {
		return Unit{"0"}
	}

	q, r := u.Measure()

	if q == 0 {
		return Unit{"0"}
	}

	switch r {
	case "px":
		return Px(q)
	case "em":
		return Em(q)
	case "rem":
		return Rem(q)
	case "vmin":
		return Vmin(q)
	case "%":
		return Percent(q)
	default:
		return Unit{"0"}
	}
}

func (u Unit) String() string {
	return u.string
}

func (Unit) unitValue()         {}
func (Unit) unitOrAutoValue()   {}
func (Unit) unitOrNoneValue()   {}
func (Unit) fontSizeValue()     {}
func (Unit) borderRadiusValue() {}
func (Unit) thicknessValue()    {}

func (u Unit) Rule() Rule {
	if u.string == "" {
		return "0"
	}
	return Rule(u.string)
}

func Em(v float64) Unit {
	return Unit{
		fmt.Sprintf(`%fem`, v),
	}
}

func Px(v float64) Unit {
	return Unit{
		fmt.Sprintf(`%fpx`, v),
	}
}

func Rem(v float64) Unit {
	return Unit{
		fmt.Sprintf(`%frem`, v),
	}
}

func Vmin(v float64) Unit {
	return Unit{
		fmt.Sprintf(`%fvmin`, v),
	}
}

func Percent(v float64) Unit {
	return Unit{
		fmt.Sprintf(`%f%%`, v),
	}
}
