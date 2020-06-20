package client

import "qlova.org/seed/js"

//String is a readonly client-typed string.
type String js.AnyString

//NewString returns a client-typed String from the given string.
func NewString(literal string) String {
	return js.NewString(literal)
}

//SideString is a variable string stored on the client.
type SideString struct {
	SideValue
}

//GetString implements String
func (s SideString) GetString() js.String {
	return js.String{s.GetValue()}
}

//Split returns an array of strings based on the split seperator.
func Split(s String, sep String) js.Array {
	return js.Array{s.GetString().Call("split", sep)}
}
