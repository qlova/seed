package client

import (
	"qlova.org/seed/js"
)

//Value is a readonly client-typed value.
type Value = js.AnyValue

//ValueOf returns a client-typed value of the given argument.
func ValueOf(any interface{}) Value {
	return js.ValueOf(any)
}
