package percentage

import "fmt"

//Int is an integer-based percentage-quantity unit.
type Int int32

func (p Int) String() string {
	return fmt.Sprintf("%v%%", int32(p))
}

//Measure implements unit.Unit
func (p Int) Measure() (quantity float64, reference string) {
	return float64(p), "%"
}

//Percentage is a percentage-quantity unit.
type Percentage float64

//New returns the given quantity as a unit.
func New(quantity int16) Percentage {
	return Percentage(quantity)
}

func (p Percentage) String() string {
	return fmt.Sprintf("%f%%", p)
}

//Measure implements unit.Unit
func (p Percentage) Measure() (quantity float64, reference string) {
	return float64(p), "%"
}
