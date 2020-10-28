package the

import "qlova.org/seed/client"

//BoolExpression is an expression containing a client.Bool.
type BoolExpression struct {
	client.Bool
	components []client.Value
}

//Bool returns a BoolExpression containing b.
func Bool(b client.Bool, components ...client.Value) BoolExpression {
	return BoolExpression{b, client.FlattenComponents(append(components, b)...)}
}

//Components implements clientside.Compound
func (b BoolExpression) Components() []client.Value {
	return b.components
}
