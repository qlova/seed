package js

import (
	"sort"
	"strconv"
	"strings"
)

//New is equivalant to calling 'new Class(args...)' in js.
func New(class AnyValue, args ...AnyValue) Value {
	return NewValue(`new %v`, Function{class.GetValue()}.Call(args...))
}

type NewObject map[string]AnyValue

func (literal NewObject) GetObject() Object {
	var object strings.Builder
	object.WriteByte('{')

	//Deterministic render.
	keys := make([]string, 0, len(literal))
	for i := range literal {
		keys = append(keys, string(i))
	}
	sort.Strings(keys)

	var i = 0

	for _, key := range keys {
		value := literal[key]

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

//Set a property of this object.
func (o Object) Set(property AnyString, value AnyValue) Script {
	if value == nil {
		value = NewValue("null")
	}
	return func(q Ctx) {
		q(o)
		q('[')
		q(property)
		q("] = ")
		q(value.GetValue())
		q(';')
	}
}

//Get gets the JavaScript property p of value v.
func (o Object) Get(property AnyString) Value {
	return NewValue(`%v[%v]`, o, property)
}
