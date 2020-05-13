package js

import (
	"strconv"
	"strings"
)

type NewObject map[string]AnyValue

func (literal NewObject) GetObject() Object {
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

func (literal NewObject) GetValue() Value {
	return literal.GetObject().Value
}

func (literal NewObject) GetBool() Bool {
	return literal.GetValue().GetBool()
}

type Object struct {
	Value
}

//AnyObject is anything that can retrieve a string.
type AnyObject interface {
	AnyValue
	GetObject() Object
}

//GetObject impliments AnyObject.
func (o Object) GetObject() Object {
	return o
}
