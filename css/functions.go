package css

import (
	"fmt"
	color_go "image/color"
	"time"
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
	var r, g, b, a = c.RGBA()
	if a != 255 {
		c := fmt.Sprint("rgba(", (float64(r)/65535)*255, ",", (float64(g)/65535)*255, ",", (float64(b)/65535)*255, ",", float64(a)/65535, ")")
		return Rule(c)
	} else {
		return Rule(fmt.Sprint("rgb(", r, ",", g, ",", b, ")"))
	}
}

type Unit struct {
	string
}

func (Unit) unitValue()         {}
func (Unit) unitOrAutoValue()   {}
func (Unit) unitOrNoneValue()   {}
func (Unit) fontSizeValue()     {}
func (Unit) borderRadiusValue() {}

func (u Unit) Rule() Rule {
	return Rule(u.string)
}

func Em(v float32) Unit {
	return Unit{
		fmt.Sprintf(`%fem`, v),
	}
}

func Rem(v float32) Unit {
	return Unit{
		fmt.Sprintf(`%frem`, v),
	}
}

func Vmin(v float32) Unit {
	return Unit{
		fmt.Sprintf(`%fvmin`, v),
	}
}

func Percent(v float32) Unit {
	return Unit{
		fmt.Sprintf(`%f%%`, v),
	}
}
