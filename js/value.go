package js

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//AnyValue is anything that can produce a js value.
type AnyValue interface {
	GetValue() Value
}

//Value is any js value.
type Value struct {
	string
}

//NewValue allows you to create a js value from a raw string.
func NewValue(format string, args ...AnyValue) Value {
	if len(args) == 0 {
		return Value{format}
	}

	var values = make([]interface{}, len(args))
	for i := range args {
		values[i] = args[i].GetValue()
	}

	var s = fmt.Sprintf(format, values...)
	return Value{s}
}

func ValueOf(literal interface{}) Value {
	b, err := json.Marshal(literal)
	if err != nil {
		panic(fmt.Errorf("js.ValueOf invalid type: %v (%w)", reflect.TypeOf(literal), err))
	}
	return NewValue(string(b))
}

//String returns the raw value.
func (v Value) String() string {
	return v.string
}

//GetValue impliments AnyValue.
func (v Value) GetValue() Value {
	return v
}

func (v Value) Var(q Ctx) Value {
	var old = v.string
	var s = q.Unique()
	v.string = s
	q("let ")
	q(v.string)
	q('=')
	q(old)
	q(';')
	return v
}

//Set sets the JavaScript property p of value v to ValueOf(x).
func (v Value) Set(property AnyString, value Value) Script {
	return func(q Ctx) {
		q(v)
		q('[')
		q(property)
		q("] = ")
		q(value)
		q(';')
	}
}
