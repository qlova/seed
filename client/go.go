package client

import (
	"encoding/base64"
	"math/big"
	"reflect"

	"qlova.org/seed/use/js"
	"qlova.org/seed/use/js/console"
)

var goExports = make(map[string]reflect.Value)

//Go requests the client to call the given Go function in a new goroutine, with the given client Values automatically converted to equivalent Go values and are passed to the given function.
//The function can optionally take a Ctx as the first argument, if so, then it is passed to the function and arguments are assigned to the following arguments.
func Go(fn interface{}, args ...Value) Script {
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

		var CallingString = `/go/` + name

		var formdata = Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		var f = js.Function{js.NewValue(formdata + `.set`)}

		if len(args) > 0 {
			for i, arg := range args {
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

		q([]byte(`seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active);`))
	})
}

//Run runs a go function, blocking until it completes.
func Run(fn interface{}, args ...Value) Script {
	return js.Script(func(q js.Ctx) {
		//Get a unique string reference for f.
		var name = base64.RawURLEncoding.EncodeToString(big.NewInt(rpcID).Bytes())

		rpcID++

		var value = reflect.ValueOf(fn)

		if value.Kind() != reflect.Func || value.Type().NumOut() > 2 {
			panic("script.Run: Must pass a Go function without zero or one return values, not a " + reflect.TypeOf(fn).String())
		}
		if value.Type().NumOut() > 2 && value.Type().Out(1) != reflect.TypeOf(error(nil)) {
			panic("script.Run: Must pass a Go function with an error value as the second parameter " + reflect.TypeOf(fn).String())
		}

		goExports[name] = value

		var CallingString = `/go/` + name

		var formdata = Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		var f = js.Function{js.NewValue(formdata + `.set`)}

		if len(args) > 0 {
			for i, arg := range args {
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

		q([]byte(`await seed.request("POST", ` + formdata + `, "` + CallingString + `", false, seed.active);`))
	})
}

//Call calls a go function and returns the result.
func Call(fn interface{}, args ...Value) Value {
	return js.Await(js.Call(js.NewFunction(func(q js.Ctx) {
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

		var CallingString = `/go/` + name

		var formdata = Unique()

		q(`let ` + formdata + ` = new FormData();`)

		//Get all positional arguments and add them to the formdata.
		var f = js.Function{js.NewValue(formdata + `.set`)}

		if len(args) > 0 {
			for i, arg := range args {
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

		q([]byte(`return await seed.request("POST", ` + formdata + `, "` + CallingString + `", false);`))
	})))
}

var rpcID int64 = 1

func Download(fn interface{}, args ...Value) Script {

	if url, ok := fn.(String); ok {
		return js.Func("c.download").Run(NewString(""), url)
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

		if len(args) > 0 {
			for i, arg := range args {
				switch arg.(type) {
				case js.AnySet:
					q.Run(f, q.String(string('a'+rune(i))), js.NewValue(`JSON.stringify(Array.from(%v))`, arg))
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
	})
}
