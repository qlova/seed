package client

import (
	"encoding/base64"
	"math/big"
	"reflect"

	"qlova.org/seed/js"
	"qlova.org/seed/js/console"
	"qlova.org/seed/script"
	"qlova.org/seed/user"
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
	var values = make([]js.AnyValue, len(args))
	for i, arg := range args {
		values[i] = arg
	}
	return script.Go(fn, values...)
}

var rpcID int64

func Download(fn interface{}, args ...Value) Script {
	return func(q script.Ctx) {
		//Get a unique string reference for f.
		var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

		rpcID++

		var value = reflect.ValueOf(fn)

		if value.Kind() != reflect.Func || value.Type().NumOut() > 2 {
			panic("script.Go: Must pass a Go function without zero or one return values, not a " + reflect.TypeOf(fn).String())
		}
		if value.Type().NumOut() > 2 && value.Type().Out(1) != reflect.TypeOf(error(nil)) {
			panic("script.Go: Must pass a Go function with an error value as the second parameter " + reflect.TypeOf(fn).String())
		}

		script.Exports[name] = value

		var CallingString = `/call/` + name + `?`

		var formdata = script.Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		var f = js.Function{js.NewValue(formdata + `.set`)}

		if len(args) > 0 {
			for i, arg := range args {
				switch arg.(type) {
				default:
					q.Run(f, q.String(string('a'+rune(i))), js.NewValue(`JSON.stringify(%v)`, arg))
				}
			}
		}

		q(console.Log(NewString(CallingString).GetString().Plus(
			js.String{js.NewValue(`new URLSearchParams(%v).toString()`, js.NewValue(formdata))},
		)))

		q(js.Func(`seed.download`).Run(NewString(""), NewString(CallingString).GetString().Plus(
			js.String{js.NewValue(`new URLSearchParams(%v).toString()`, js.NewValue(formdata))},
		)))
	}
}
