package client

import (
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user"
)

//Ctx is a client ctx
type Ctx = user.Ctx

//Script instructs the client to do something.
type Script = js.Script

//If runs the provided scripts if the clients condition is true.
func If(condition Bool, do ...Script) Script {
	return js.If(condition, script.New(do...))
}

//Go requests the client to call the given Go function with the given client Values automatically converted to equivalent Go values and are passed to the given function.
//The function can optionally take a Ctx as the first argument, if so, then it is passed to the function and arguments are assigned to the following arguments.
func Go(fn interface{}, args ...Value) Script {
	return script.Go(fn, args...)
}
