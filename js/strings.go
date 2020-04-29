package js

import "strconv"

//String is a javascript string.
type String struct {
	Value
}

//AnyString is anything that can retrieve a string.
type AnyString interface {
	AnyValue
	GetString() String
}

//NewString returns a new javascript string from a Go literal.
func NewString(literal string) String {
	return String{NewValue(strconv.Quote(literal))}
}

//String is shorthand for NewString.
func (Ctx) String(literal string) String {
	return NewString(literal)
}

//GetString impliments AnyString.
func (s String) GetString() String {
	return s
}

//Equals returns true if the two strings are equal.
func (s String) Equals(b AnyString) Bool {
	return Bool{NewValue("(" + s.string + "==" + b.GetString().string + ")")}
}
