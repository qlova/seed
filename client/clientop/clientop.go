package clientop

import (
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/js"
)

//Bool is a client bool.
type Bool struct {
	client.Bool
	components []client.Value
}

func NewBool(b client.Bool, components ...client.Value) Bool {
	return Bool{b, client.FlattenComponents(components)}
}

//Components implements clientside.Compound
func (b Bool) Components() []client.Value {
	return b.components
}

//Eq returns true if both values are equal.
func Eq(a, b client.Value) Bool {
	if seca, ok := a.(*clientside.Secret); ok {
		if sb, ok := b.(client.String); ok {
			return NewBool(seca.Equals(sb), a, b)
		}
	}
	return NewBool(js.NewValue("(%v == %v)", a, b), a, b)
}

//NotEq returns true if both values are different.
func NotEq(a, b client.Value) Bool {
	if seca, ok := a.(*clientside.Secret); ok {
		if sb, ok := b.(client.String); ok {
			return Not(NewBool(seca.Equals(sb), a, b))
		}
	}
	return NewBool(js.NewValue("(%v != %v)", a, b), a, b)
}

//Not returns the inverse of the provided bool.
func Not(a client.Bool) Bool {
	return NewBool(js.NewValue("(!%v)", a), a)
}

//And returns true if all values are true.
func And(a, b client.Value, more ...client.Value) Bool {
	if len(more) == 0 {
		return NewBool(js.NewValue("(%v && %v)", a, b), a, b)
	}

	var components = []client.Value{a, b}

	var expression = "(%v && %v"

	for _, v := range more {
		expression += " && %v"
		components = append(components, v)
	}

	expression += ")"

	return NewBool(js.NewValue(expression, components...), components...)
}

//Or returns true if one of the values is true.
func Or(a, b client.Value, more ...client.Value) Bool {
	if len(more) == 0 {
		return NewBool(js.NewValue("(%v || %v)", a, b), a, b)
	}

	var components = []client.Value{a, b}

	var expression = "(%v || %v"

	for _, v := range more {
		expression += " || %v"
		components = append(components, v)
	}

	expression += ")"

	return NewBool(js.NewValue(expression, components...), components...)
}

//Number is a client number.
type Number struct {
	js.Number
	components []client.Value
}

var _ client.Value = Number{}

func NewNumber(n js.Number, components ...client.Value) Number {
	return Number{n, client.FlattenComponents(components)}
}

//Components implements clientside.Compound
func (n Number) Components() []client.Value {
	return n.components
}

//Add returns a+b
func Add(a, b js.AnyNumber) Number {
	return NewNumber(js.Number{js.NewValue("(%v + %v)", a, b)}, a, b)
}

//Divide returns a/b
func Divide(a, b js.AnyNumber) Number {
	return NewNumber(js.Number{js.NewValue("(%v / %v)", a, b)}, a, b)
}

//Subtract returns a-b
func Subtract(a, b js.AnyNumber) Number {
	return NewNumber(js.Number{js.NewValue("(%v - %v)", a, b)}, a, b)
}

func LessThanEqual(a, b js.AnyNumber) Bool {
	return NewBool(js.Number{js.NewValue("(%v <= %v)", a, b)}, a, b)
}

//GreaterThan returns a>b
func GreaterThan(a, b js.AnyNumber) Bool {
	return NewBool(js.Number{Value: js.NewValue("(%v > %v)", a, b)}, a, b)
}
