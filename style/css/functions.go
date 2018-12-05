package css

import (
	"fmt"
	go_color "image/color"
)

const Em = 1i
const Px = 1i + 1

func Decode(unit complex128) unitType {
	if unit == 0 {
		return "initial"
	}

	if imag(unit) == real(unit) {
		return unitType(fmt.Sprint(real(unit), "px"))
	}

	if imag(unit) != 0 {
		return unitType(fmt.Sprint(imag(unit), "em"))
	} else {
		return unitType(fmt.Sprint(real(unit), "%"))
	}
}

func AnimationName(name string) animationNameValue {
	return animationNameType(name)
}

//Shorthand for Percent()
func (n numberType) Pc() unitType {
	return n.Percent()
}

//Returns Percentage units.
func (n numberType) Percent() unitType {
	return unitType(n.String() + "%")
}

//Shorthand for Pixels()
func (n numberType) Px() unitType {
	return unitType(n.String() + "px")
}

//Returns Pixel units.
func (n numberType) Pixels() unitType {
	return unitType(n.String() + "px")
}

//Returns Pixel units.
func (n numberType) Em() unitType {
	return unitType(n.String() + "em")
}

//Returns ViewportWidth units.
func (n numberType) Vw() unitType {
	return unitType(n.String() + "vw")
}

//Returns ViewportHeight units.
func (n numberType) Vh() unitType {
	return unitType(n.String() + "vh")
}

//Returns a CSS number.
func Number(number float64) numberType {
	return numberType(fmt.Sprint(number))
}

//Returns a CSS number.
func Integer(integer int) integerType {
	return integerType(fmt.Sprint(integer))
}

//Returns a CSS time.
func Time(seconds float64) timeType {
	return timeType(fmt.Sprint(seconds, "s"))
}

//Returns a CSS color.
func Colour(c go_color.Color) colorType {
	var r, g, b, a = c.RGBA()
	return colorType(fmt.Sprint("rgba(", r, ",", g, ",", b, ",", a, ")"))
}

func LinearGradient(direction float64, start, end colorValue) gradientType {
	return gradientType(fmt.Sprint("linear-gradient(", direction, "rad,", start, ",", end, ")"))
}

var Linear animationTimingFunctionType = "linear"
var EaseInOut animationTimingFunctionType = "ease-in-out"




