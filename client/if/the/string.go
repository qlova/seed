package the

import (
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

//StringExpression is an expression containing a client.String.
type StringExpression struct {
	client.String
	components []client.Value
}

//String returns a StringExpression containing s.
func String(s client.String, components ...client.Value) StringExpression {
	return StringExpression{s, client.FlattenComponents(append(components, s)...)}
}

//Components implements clientside.Compound
func (a StringExpression) Components() []client.Value {
	return a.components
}

//Is returns a == b as a BoolExpression
func (a StringExpression) Is(literal string) BoolExpression {
	b := client.NewString(literal)
	return Bool(js.NewValue("(%v == %v)", a, b), a, b)
}

//TernaryString is a StringExpression that can be one of two values depending on a condition.
type TernaryString struct {
	Condition client.Bool
	IfTrue    StringExpression
}

//If returns a TernaryString.
func (a StringExpression) If(condition client.Bool) TernaryString {
	return TernaryString{condition, a}
}

//Otherwise returns the complete StringExpression of the TernaryString.
func (t TernaryString) Otherwise(value client.String) StringExpression {
	return String(js.String{Value: js.NewValue("(%v ? %v : %v)", t.Condition, t.IfTrue, value)}, t.Condition, t.IfTrue, value)
}
