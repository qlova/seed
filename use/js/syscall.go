package js

//Global returns the JavaScript global object, usually "window" or "global".
func Global() Value {
	return NewValue("(window || global)")
}

//Null returns the JavaScript value "null".
func Null() Value {
	return NewValue("null")
}

//Undefined returns the JavaScript value "undefined".
func Undefined() Value {
	return NewValue("undefined")
}
