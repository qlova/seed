package ratio

import "fmt"

//Ratio is a percentage-quantity unit.
type Ratio float64

//New returns the given quantity as a unit.
func New(quantity float64) Ratio {
	return Ratio(quantity)
}

func (r Ratio) String() string {
	return fmt.Sprintf("%f%%", r)
}

//Measure implements unit.Unit
func (r Ratio) Measure() (quantity float64, reference string) {
	return float64(r), "%"
}
