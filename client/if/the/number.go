package the

import (
	"qlova.org/seed/client"
	"qlova.org/seed/js"
)

//NumberExpression is an expression containing a number.
type NumberExpression struct {
	client.Float

	components []client.Value
}

//Components implements clientside.Compound
func (a NumberExpression) Components() []client.Value {
	return a.components
}

//IsLessThan returns a < b
func (a NumberExpression) IsLessThan(b client.Float) BoolExpression {
	return Bool(js.NewValue("(%v < %v)", a, b), a, b)
}

//IsMoreThan returns a > b
func (a NumberExpression) IsMoreThan(b client.Float) BoolExpression {
	return Bool(js.NewValue("(%v > %v)", a, b), a, b)
}

//Number returns the NumberType of n.
func Number(n client.Float, components ...client.Value) NumberExpression {
	return NumberExpression{n, client.FlattenComponents(append(components, n)...)}
}

//SumOf returns the sum of (a + b + others...)
func SumOf(a, b client.Float, others ...client.Float) NumberExpression {
	var expression string = "(%v + %v"
	var values = []client.Value{a, b}

	for _, n := range others {
		expression += " + %v"
		values = append(values, n)
	}

	expression += ")"

	return Number(js.Number{Value: js.NewValue(expression, values...)}, values...)
}

//ProductOf returns the product of (a * b * others...)
func ProductOf(a, b client.Float, others ...client.Float) NumberExpression {
	var expression string = "(%v * %v"
	var values = []client.Value{a, b}

	for _, n := range others {
		expression += " * %v"
		values = append(values, n)
	}

	expression += ")"

	return Number(js.Number{Value: js.NewValue(expression, values...)}, values...)
}

//QuotientOf returns the quotient of (a / b / others...)
func QuotientOf(a, b client.Float, others ...client.Float) NumberExpression {
	var expression string = "(%v / %v"
	var values = []client.Value{a, b}

	for _, n := range others {
		expression += " / %v"
		values = append(values, n)
	}

	expression += ")"

	return Number(js.Number{Value: js.NewValue(expression, values...)}, values...)
}

//DifferenceOf returns the quotient of (a - b - others...)
func DifferenceOf(a, b client.Float, others ...client.Float) NumberExpression {
	var expression string = "(%v - %v"
	var values = []client.Value{a, b}

	for _, n := range others {
		expression += " - %v"
		values = append(values, n)
	}

	expression += ")"

	return Number(js.Number{Value: js.NewValue(expression, values...)}, values...)
}
