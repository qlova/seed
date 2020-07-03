package js

import "fmt"

//Number is a javascript number.
type Number struct {
	Value
}

//AnyNumber is anything that can retrieve a number.
type AnyNumber interface {
	AnyValue
	GetNumber() Number
}

//NewNumber returns a new javascript number from a Go literal.
func NewNumber(literal float64) Number {
	return Number{NewValue(fmt.Sprintf("%#v", literal))}
}

//Number is shorthand for NewNumber.
func (Ctx) Number(literal float64) Number {
	return NewNumber(literal)
}

//GetNumber impliments AnyNumber.
func (n Number) GetNumber() Number {
	return n
}

//Plus returns the two numbers added together.
func (n Number) Plus(b AnyNumber) Number {
	return Number{NewValue(fmt.Sprintf(`(%v+%v)`, n, b.GetNumber()))}
}

//Minus returns the two numbers subtracted together.
func (n Number) Minus(b AnyNumber) Number {
	return Number{NewValue(fmt.Sprintf(`(%v-%v)`, n, b.GetNumber()))}
}

//ToString returns the number as a string.
func (n Number) ToString() String {
	return String{n.Call("toString")}
}

//DivideBy returns n/b
func (n Number) DivideBy(b AnyNumber) Number {
	return Number{NewValue(fmt.Sprintf(`(%v/%v)`, n, b.GetNumber()))}
}

//ModBy returns n%b
func (n Number) ModBy(b AnyNumber) Number {
	return Number{NewValue(fmt.Sprintf(`(%v%%%v)`, n, b.GetNumber()))}
}
