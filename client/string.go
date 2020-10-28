package client

import "qlova.org/seed/web/js"

//String is a readonly client-typed string.
type String js.AnyString

//NewString returns a client-typed String from the given string.
func NewString(literal string) String {
	return js.NewString(literal)
}

//Split returns an array of strings based on the split seperator.
func Split(s String, sep String) js.Array {
	return js.Array{Value: s.GetString().Call("split", sep)}
}
