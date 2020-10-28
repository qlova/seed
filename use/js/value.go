package js

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

//AnyValue is anything that can produce a js value.
type AnyValue interface {
	GetValue() Value
	GetBool() Bool
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

//GetBool impliments AnyBool.
func (v Value) GetBool() Bool {
	v.string = `(!!` + v.string + ")"
	return Bool{v}
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
func (v Value) Set(property string, value AnyValue) Script {
	if value == nil {
		value = NewValue("null")
	}
	return func(q Ctx) {
		q(v)
		q('[')
		q(strconv.Quote(property))
		q("] = ")
		q(value.GetValue())
		q(';')
	}
}

//Get gets the JavaScript property p of value v.
func (v Value) Get(property string) Value {
	v.string = v.string + "[" + strconv.Quote(property) + "]"
	return v
}

//Index gets the JavaScript index i of value v.
func (v Value) Index(i int) Value {
	v.string = v.string + "[" + strconv.Itoa(i) + "]"
	return v
}

//Call calls the method on the given value.
func (v Value) Call(method string, args ...AnyValue) Value {
	return Call(Function{NewValue(v.string + "." + method)}, args...)
}

//Run runs the method on the given value.
func (v Value) Run(method string, args ...AnyValue) Script {
	return Run(Function{NewValue(v.string + "." + method)}, args...)
}

//Equals returns true if the two value are equal.
func (v Value) Equals(other Value) Bool {
	return NewValue(`(%v == %v)`, v, other).GetBool()
}
