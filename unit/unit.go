package unit

//Unit is a type of unit.
type Unit interface {
	//String returns a string representation of the unit.
	String() string

	//Unit returns both the quantity and a unit reference.
	Measure() (quantity float64, reference string)
}
