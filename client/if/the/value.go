package the

import (
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/js"
)

//Value is a value.
type Value struct {
	client.Value

	components []client.Value
}

//Components implements clientside.Compound
func (a Value) Components() []client.Value {
	return a.components
}

//ValueOf returns the Value of v.
func ValueOf(v client.Value, components ...client.Value) Value {
	return Value{v, client.FlattenComponents(append(components, v)...)}
}

//Is returns a == b as a Value
func (a Value) Is(b client.Value) BoolExpression {

	//Special optimised case.
	if seca, ok := a.Value.(*clientside.Secret); ok {
		if sb, ok := b.(client.String); ok {
			return Bool(seca.Equals(sb), a, b)
		}
	}

	return Bool(js.NewValue("(%v == %v)", a, b), a, b)
}

//IsNot returns a != b as a Value
func (a Value) IsNot(b client.Value) BoolExpression {

	//Special optimised case.
	if seca, ok := a.Value.(*clientside.Secret); ok {
		if sb, ok := b.(client.String); ok {
			return Bool(seca.Equals(sb).GetBool().Not(), a, b)
		}
	}

	return Bool(js.NewValue("(%v != %v)", a, b), a, b)
}
