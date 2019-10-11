//+build wasm

package js

import (
	"syscall/js"

	"github.com/qlova/seed"
)

//Value is a generic js.Value.
type Value struct {
	js.Value
}

//ValueFromUpdate returns a value from a seed.Update
func ValueFromUpdate(u seed.Update) Value {
	return Value{
		js.Global().Get("document").Call("getElementById", u.ID()),
	}
}

func (v Value) Global() Value {
	return Value{js.Global()}
}

func (v Value) String(s string) Value {
	return Value{js.ValueOf(s)}
}

func (v Value) Number(n float64) Value {
	return Value{js.ValueOf(n)}
}

func (v Value) Run(method string, args ...Value) {
	var in = make([]interface{}, len(args))
	for i := range args {
		in[i] = args[i]
	}
	v.Value.Call(method, in...)
}

func (v Value) Call(method string, args ...Value) Value {
	var in = make([]interface{}, len(args))
	for i := range args {
		in[i] = args[i]
	}
	return Value{v.Value.Call(method, in...)}
}

func (v Value) Set(property string, value Value) {
	v.Value.Set(property, value)
}

func (v Value) Get(property string) Value {
	return Value{v.Value.Get(property)}
}

func (v Value) Var(name string) Value {
	return v
}
