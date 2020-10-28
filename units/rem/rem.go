package rem

import "fmt"

//Unit is equal to the global font-size setting of the user.
type Unit float64

//One rem.
const One = Unit(1)

//New returns the given quantity as a unit.
func New(quantity float64) Unit {
	return Unit(quantity)
}

func (u Unit) String() string {
	return fmt.Sprintf("%frem", u)
}

//Measure implements unit.Unit
func (u Unit) Measure() (quantity float64, reference string) {
	return float64(u), "rem"
}
