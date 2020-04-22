package js

import (
	"strconv"
	"strings"
)

type Object struct {
	Value
}

//AnyObject is anything that can retrieve a string.
type AnyObject interface {
	AnyValue
	GetObject() Object
}

//NewObject returns a new javascript string from a Go literal.
func NewObject(literal map[string]AnyValue) Object {
	var object strings.Builder
	object.WriteByte('{')
	var i = 0
	for key, value := range literal {
		object.WriteString(strconv.Quote(key))
		object.WriteByte(':')
		object.WriteString(value.GetValue().String())
		if i < len(literal)-1 {
			object.WriteByte(',')
		}
		i++
	}
	object.WriteByte('}')

	return Object{NewValue(object.String())}
}

//Object is shorthand for NewObject.
func (Ctx) Object(literal map[string]AnyValue) Object {
	return NewObject(literal)
}

//GetObject impliments AnyObject.
func (o Object) GetObject() Object {
	return o
}
