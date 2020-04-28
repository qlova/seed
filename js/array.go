package js

import (
	"strings"
)

type Array struct {
	Value
}

//AnyArray is anything that can retrieve a string.
type AnyArray interface {
	AnyValue
	GetArray() Array
}

//NewArray returns a new javascript string from a Go literal.
func NewArray(literal []AnyValue) Array {
	var object strings.Builder

	object.WriteByte('[')
	var i = 0
	for _, value := range literal {
		object.WriteString(value.GetValue().String())
		if i < len(literal)-1 {
			object.WriteByte(',')
		}
		i++
	}
	object.WriteByte(']')

	return Array{NewValue(object.String())}
}

//Array is shorthand for NewArray.
func (Ctx) Array(literal []AnyValue) Array {
	return NewArray(literal)
}

//GetArray impliments AnyArray.
func (a Array) GetArray() Array {
	return a
}

//Index returns the value inside the array at the given index.
func (a Array) Index(i AnyNumber) Value {
	a.string += "[" + i.GetNumber().String() + "]"
	return a.Value
}
