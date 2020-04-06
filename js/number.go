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
	return Number{Value{fmt.Sprintf("%#v", literal)}}
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
func (n Number) Plus(b Number) Number {
	return Number{Value{fmt.Sprintf(`(%v+%v)`, n, b)}}
}
