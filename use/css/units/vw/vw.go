package vw

import "fmt"

//Unit is equal to the 1/100ths of the shortest viewport dimension.
type Unit float64

//New returns the given quantity as a unit.
func New(quantity float64) Unit {
	return Unit(quantity)
}

func (u Unit) String() string {
	return fmt.Sprintf("%fvw", u)
}

//Measure implements unit.Unit
func (u Unit) Measure() (quantity float64, reference string) {
	return float64(u), "vw"
}
