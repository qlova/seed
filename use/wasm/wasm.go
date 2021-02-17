// +build wasm

package wasm

import (
	"bytes"
	"errors"
	"fmt"
	"io"
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

func await(v js.Value) (result js.Value, ok bool) {
	if v.Type() != js.TypeObject || v.Get("then").Type() != js.TypeFunction {
		return v, true
	}

	done := make(chan struct{})

	onResolve := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result = args[0]
		ok = true
		close(done)
		return nil
	})
	defer onResolve.Release()

	onReject := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result = args[0]
		ok = false
		close(done)
		return nil
	})
	defer onReject.Release()

	v.Call("then", onResolve, onReject)
	<-done
	return
}

type blobReader struct {
	index int

	blob js.Value
}

func (b *blobReader) Read(bytes []byte) (int, error) {
	if b.blob.IsUndefined() {
		return 0, errors.New("invalid blob")
	}

	if b.index >= b.blob.Get("size").Int() {
		return 0, io.EOF
	}

	arraybuffer, _ := await(b.blob.Call("slice", b.index, b.index+len(bytes)).Call("arrayBuffer"))

	b.index += len(bytes)

	i := js.CopyBytesToGo(bytes, js.Global().Get("Uint8Array").New(arraybuffer))

	return i, nil
}

var _ io.Reader = new(blobReader)

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

			case js.TypeObject:
				if jsv.InstanceOf(js.Global().Get("Blob")) {
					val = &blobReader{0, jsv}
				} else {
					panic("unsupported object type")
				}

			default:
				val = js.Global().Get("JSON").Call("stringify", jsv).String()
			}

			iargs = append(iargs, val)
		}

		var resolve, reject js.Value

		var promise = js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve = args[0]
			reject = args[1]
			return nil
		}))

		go func() {
			i, err := ctx.Call(f, iargs...)
			if err != nil {
				log.Println(err)
			}

			ctx.Return(i, err)

			v, _ := await(js.Global().Get("AsyncFunction").New(buffer.String()).Invoke())

			resolve.Invoke(v)
		}()

		return promise
	}))
}
