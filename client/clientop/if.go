package clientop

import (
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientfmt"
	"qlova.org/seed/js"
)

//Value is a client value.
type Value struct {
	client.Value
	components []client.Value
}

func NewValue(v client.Value, components ...client.Value) Value {
	return Value{v, client.FlattenComponents(components)}
}

//Components implements clientside.Compound
func (v Value) Components() []client.Value {
	return v.components
}

func (v Value) String() clientfmt.String {
	return clientfmt.NewString(js.String{Value: v.Value.GetValue()}, v.components...)
}

type Ternary struct {
	condition client.Bool
	value     client.Value
}

func (t Ternary) Else(value client.Value) Value {
	return NewValue(js.NewValue("(%v ? %v : %v)", t.condition, t.value, value), t.condition, t.value, value)
}

func If(condition client.Bool, value client.Value) Ternary {
	return Ternary{
		condition: condition,
		value:     value,
	}
}
