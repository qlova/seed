//Package unit provides useful Qlovaseed constants.
package unit

import (
	"math"

	"github.com/qlova/seed/style/css"
)

//Unit type.
type Unit = complex128

//Auto is a an automatic unit that works in some contexts.
const Auto = math.MaxFloat64

//Em is the 'default font size' unit.
//This is the recommended unit to use in almost all circumstances.
const Em = css.Em

//Vmin is equal to 1% of the viewport's smaller dimension
const Vmin = css.Vm

//In is an inch.
const In = 96 * Px

//Px are device-independent pixels.
const Px = css.Px

//Mm is a millimeter.
const Mm = In * 0.0393701

//Cm is a centimeter.
const Cm = 10 * Mm

//Pt is a CSS pt.
const Pt = In / 72

//Pc is a pica.
const Pc = 12 * Pt

//Deg returns degrees to radians.
func Deg(deg float64) float64 {
	return deg * math.Pi / 180
}
