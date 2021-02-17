package client

import (
	"encoding/base64"
	"math/big"
	"reflect"

	"qlova.org/seed/use/js"
	"qlova.org/seed/use/wasm"
)

var goExports = make(map[string]reflect.Value)

//MutableFloat is a float that can be mutated.
type MutableFloat interface {
	Float

	SetTo(Float) Script
}

type recording struct {
	Function interface{}
	Progress MutableFloat
}

//Record can be used to record the progress of Go and Run calls.
func Record(fn interface{}, progress MutableFloat) interface{} {
	return recording{
		fn, progress,
	}
}

//Argument types can decide how to encode themselves as arguments to a Go or Run call.
type Argument interface {
	Value

	AsArgument() Value
}

//Parser is used to parse a Argument.
type Parser interface {
	Parse(string) error
}

//Validator is used to validate an Argument, if the validation fails, the request returns the provided error.
type Validator interface {
	Validate() error
}

func rpc(q js.Ctx, fn interface{}, args ...Value) (progress recording, CallingString, formdata string) {
	rec, recording := fn.(recording)
	if recording {
		fn = rec.Function
		progress = rec
	}

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

	goExports[name] = value

	CallingString = `/go/` + name

	formdata = Unique()

	q(`let ` + formdata + ` = new FormData();`)

	//Get all positional arguments and add them to the formdata.
	var f = js.Function{js.NewValue(formdata + `.set`)}

	if len(args) > 0 {
		for i, arg := range args {
			if a, ok := arg.(Argument); ok {
				arg = a.AsArgument()
			}

			switch arg.(type) {
			case File:
				q.Run(f, q.String(string('a'+rune(i))), arg)
			case js.AnySet:
				q.Run(f, q.String(string('a'+rune(i))), js.NewValue(`JSON.stringify(Array.from(%v))`, arg))
			default:
				q.Run(f, q.String(string('a'+rune(i))), js.NewValue(`JSON.stringify(%v)`, arg))
			}
		}
	}

	return
}

//Go requests the client to call the given Go function in a new goroutine, with the given client Values automatically converted to equivalent Go values and are passed to the given function.
//The function can optionally take a Ctx as the first argument, if so, then it is passed to the function and arguments are assigned to the following arguments.
func Go(fn interface{}, args ...Value) Script {
	return js.Script(func(q js.Ctx) {
		rec, CallingString, formdata := rpc(q, fn, args...)

		if rec.Function != nil {
			q([]byte(`seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active, async function(progress) {`))
			q(rec.Progress.SetTo(js.Number{Value: js.NewValue("progress")}))
			q("});")
		} else {
			q([]byte(`seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active);`))
		}
	})
}

//Run runs a go function, blocking until it completes.
func Run(fn interface{}, args ...Value) Script {
	for i, arg := range args {
		if a, ok := arg.(Argument); ok {
			args[i] = a.AsArgument()
		}
	}

	if wasm.Exported(fn) {
		return wasm.Run(fn, args...)
	}

	return js.Script(func(q js.Ctx) {
		rec, CallingString, formdata := rpc(q, fn, args...)

		if rec.Function != nil {
			q([]byte(`await seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active, async function(progress) {`))
			q(rec.Progress.SetTo(js.Number{Value: js.NewValue("progress")}))
			q("});")
		} else {
			q([]byte(`await seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active);`))
		}
	})
}

//Call calls a go function and returns the result.
func Call(fn interface{}, args ...Value) Value {
	for i, arg := range args {
		if a, ok := arg.(Argument); ok {
			args[i] = a.AsArgument()
		}
	}

	if wasm.Exported(fn) {
		return wasm.Call(fn, args...)
	}

	return js.Await(js.Call(js.NewFunction(func(q js.Ctx) {
		rec, CallingString, formdata := rpc(q, fn, args...)

		if rec.Function != nil {
			q([]byte(`return await seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active, async function(progress) {`))
			q(rec.Progress.SetTo(js.Number{Value: js.NewValue("progress")}))
			q("});")
		} else {
			q([]byte(`return await seed.request("POST", ` + formdata + `, "` + CallingString + `", false);`))
		}
	})))
}

var rpcID int64 = 1

type Name struct {
	String
}

func NameAs(s String) Name {
	return Name{s}
}

func Download(fn interface{}, args ...Value) Script {

	for i, arg := range args {
		if a, ok := arg.(Argument); ok {
			args[i] = a.AsArgument()
		}
	}

	if url, ok := fn.(String); ok {
		return js.Func("c.download").Run(NewString(""), url)
	}

	if wasm.Exported(fn) {
		return wasm.Download(fn, args...)
	}

	return js.Script(func(q js.Ctx) {
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

		goExports[name] = value

		var CallingString = `/go/` + name + `?`

		var formdata = Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		var f = js.Function{js.NewValue(formdata + `.set`)}

		var filename = NewString("")

		if len(args) > 0 {
			var skip = 0

			for i, arg := range args {
				switch v := arg.(type) {
				case js.AnySet:
					q.Run(f, q.String(string('a'+rune(i-skip))), js.NewValue(`JSON.stringify(Array.from(%v))`, arg))
				case Name:
					filename = v.String
				default:
					q.Run(f, q.String(string('a'+rune(i-skip))), js.NewValue(`JSON.stringify(%v)`, arg))
				}
			}
		}

		q(js.Func(`seed.download`).Run(filename, NewString(CallingString).GetString().Plus(
			js.String{Value: js.NewValue(`new URLSearchParams(%v).toString()`, js.NewValue(formdata))},
		)))
	})
}
