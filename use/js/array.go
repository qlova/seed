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

type NewArray []AnyValue

func (array NewArray) String() string {
	return array.GetValue().String()
}

func (literal NewArray) GetArray() Array {
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

func (literal NewArray) GetValue() Value {
	return literal.GetArray().Value
}

func (literal NewArray) GetBool() Bool {
	return literal.GetValue().GetBool()
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
