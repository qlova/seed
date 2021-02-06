// +build wasm

package wasm

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"syscall/js"

	qjs "qlova.org/seed/use/js"
)

const InstantiateStreaming = ""

func Run(f interface{}, args ...qjs.AnyValue) qjs.Script {
	return nil
}

var exports = make(map[string]struct{})

//Exported returns true if the given function was exported with this package.
func Exported(f interface{}) bool {
	_, ok := exports[runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()]
	return ok
}

func Export(f interface{}) {

	//Parser is used to parse a Argument.
	type Parser interface {
		Parse(string) error
	}

	exports[runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()] = struct{}{}
	js.Global().Set(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				debug.PrintStack()
			}
		}()

		var rvalue = reflect.ValueOf(f)
		var rtype = rvalue.Type()
		var converted = make([]reflect.Value, rtype.NumIn())

		for i := 0; i < rtype.NumIn(); i++ {

			if rtype.In(i).Implements(reflect.TypeOf([0]Parser{}).Elem()) && rtype.In(i).Kind() == reflect.Ptr {
				value := reflect.New(rtype.In(i).Elem())

				if err := value.Interface().(Parser).Parse(args[i].String()); err != nil {
					//handle(err)
					fmt.Println(err)
					return nil
				}

				converted[i] = value

			} else {

				switch rtype.In(i).Kind() {
				case reflect.String:
					converted[i] = reflect.ValueOf(args[i].String())
				}
			}
		}

		results := reflect.ValueOf(f).Call(converted)

		if len(results) == 0 {
			return nil
		}

		var result = results[0].Interface()

		switch v := result.(type) {
		case qjs.AnyScript:
			var buffer bytes.Buffer
			ctx := qjs.NewCtx(&buffer)
			ctx(v.GetScript())
			ctx.Flush()

			return js.Global().Get("AsyncFunction").New(buffer.String()).Invoke()
		}

		return nil
	}))
}
