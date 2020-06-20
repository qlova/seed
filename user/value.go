package user

import "qlova.org/seed/js"

type String js.AnyString
type Bool js.AnyBool
type Int js.AnyNumber

type Choice int32

func (c Choice) GetNumber() js.Number {
	return js.NewNumber(float64(c))
}

func (c Choice) GetValue() js.Value {
	return c.GetNumber().Value
}

func (c Choice) GetBool() js.Bool {
	return c.GetValue().GetBool()
}
