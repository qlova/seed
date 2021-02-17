// +build wasm

package wasm

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"runtime/debug"
	"syscall/js"

	"qlova.org/seed/client/clientrpc"
	qjs "qlova.org/seed/use/js"
)

const InstantiateStreaming = ""

func Run(f interface{}, args ...qjs.AnyValue) qjs.Script {
	return nil
}

func Call(f interface{}, args ...qjs.AnyValue) qjs.AnyValue {
	return nil
}

var exports = make(map[string]struct{})

//Exported returns true if the given function was exported with this package.
func Exported(f interface{}) bool {
	_, ok := exports[runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()]
	return ok
}

func Export(f interface{}) {
	exports[runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()] = struct{}{}
	js.Global().Set(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				debug.PrintStack()
			}
		}()

		var buffer bytes.Buffer

		var cr = NewRequest(&buffer)

		ctx := clientrpc.Context{Request: cr}

		var iargs []interface{}

		for i := 0; i < reflect.ValueOf(f).Type().NumIn(); i++ {
			var jsv = args[i]
			var val interface{}

			switch jsv.Type() {
			case js.TypeString:
				val = jsv.String()
			default:
				val = js.Global().Get("JSON").Call("stringify", jsv).String()
			}

			iargs = append(iargs, val)
		}

		i, err := ctx.Call(f, iargs...)
		if err != nil {
			log.Println(err)
		}

		ctx.Return(i, err)

		return js.Global().Get("AsyncFunction").New(buffer.String()).Invoke()
	}))
}
