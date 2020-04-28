package js

//Bool is a javascript boolean.
type Bool struct {
	Value
}

//AnyBool is anything that can retrieve a bool.
type AnyBool interface {
	AnyValue
	GetBool() Bool
}

//NewBool returns a new javascript boolean from a Go literal.
func NewBool(literal bool) Bool {
	if literal {
		return Bool{NewValue("true")}
	}
	return Bool{NewValue("false")}
}

//Bool is shorthand for NewBool.
func (Ctx) Bool(literal bool) Bool {
	return NewBool(literal)
}

//GetBool impliments AnyBool.
func (b Bool) GetBool() Bool {
	return b
}

//Not returns not bool.
func (b Bool) Not() Bool {
	b.Value.string = "!" + b.Value.string
	return b
}

var False = NewBool(false)
var True = NewBool(true)

//Truthy returns a boolean from any value.
func Truthy(v AnyValue) Bool {
	return Bool{Value: v.GetValue()}
}
