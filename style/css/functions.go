package css

import (
	"fmt"
	go_color "image/color"
	"math"
	"strconv"
)

const Pc = 1
const Em = 1i
const Px = 1i + 1

const Vm = 1i - 1

func Decode(unit complex128) unitType {
	if unit == 0 {
		return "0"
	}

	if real(unit) == math.MaxFloat64 {
		return "auto"
	}

	if imag(unit) == real(unit) {
		return unitType(fmt.Sprint(real(unit), "px"))
	}

	if imag(unit) == -real(unit) {
		return unitType(fmt.Sprint(imag(unit), "vmin"))
	}

	if imag(unit) != 0 {
		return unitType(fmt.Sprint(imag(unit), "rem"))
	} else {
		return unitType(fmt.Sprint(real(unit), "%"))
	}
}

func Image(url string) imageValue {
	return imageType("url(" + strconv.Quote(url) + ")")
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

//Returns ViewportMin units.
func (n numberType) Vmin() unitType {
	return unitType(n.String() + "vmin")
}

//Returns ViewportMax units.
func (n numberType) Vmax() unitType {
	return unitType(n.String() + "vmax")
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

//Returns a CSS time.
func Translate(x, y unitType) transformType {
	return transformType(fmt.Sprint("translate(", x.String(), ",", y.String(), ")"))
}

//Returns a CSS time.
func TranslateX(x unitType) transformType {
	return transformType(fmt.Sprint("translateX(", x.String(), ")"))
}

//Returns a CSS time.
func TranslateY(y unitType) transformType {
	return transformType(fmt.Sprint("translateY(", y.String(), ")"))
}

//Returns a CSS time.
func Rotate(angle float64) transformType {
	return transformType(fmt.Sprint("rotate(", angle, "rad)"))
}

//Returns a CSS time.
func RotateX(angle float64) transformType {
	return transformType(fmt.Sprint("rotateX(", angle, "rad)"))
}

//Returns a CSS time.
func Scale(x, y float64) transformType {
	return transformType(fmt.Sprint("scale(", x, ",", y, ")"))
}

//Returns a CSS time.
func Skew(x, y float64) transformType {
	return transformType(fmt.Sprint("skew(", x, "rad,", y, "rad)"))
}

//Returns a CSS color.
func Colour(c go_color.Color) colorType {
	var r, g, b, a = c.RGBA()
	if a != 255 {
		c := fmt.Sprint("rgba(", (float64(r)/65535)*255, ",", (float64(g)/65535)*255, ",", (float64(b)/65535)*255, ",", float64(a)/65535, ")")
		return colorType(c)
	} else {
		return colorType(fmt.Sprint("rgb(", r, ",", g, ",", b, ")"))
	}
}

type ColorValue = colorType

func LinearGradient(direction float64, start, end colorValue) gradientType {
	return gradientType(fmt.Sprint("linear-gradient(", direction, "rad,", start, ",", end, ")"))
}

var Linear animationTimingFunctionType = "linear"
var EaseInOut animationTimingFunctionType = "ease-in-out"
