// +build wasm

package wasm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"syscall/js"
	"time"

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

		handle := func(err error) interface{} {
			var buffer bytes.Buffer
			ctx := qjs.NewCtx(&buffer)
			ctx(qjs.Throw(qjs.NewString(err.Error())))
			ctx.Flush()
			return js.Global().Get("AsyncFunction").New(buffer.String()).Invoke()
		}

		var rvalue = reflect.ValueOf(f)
		var rtype = rvalue.Type()
		var converted = make([]reflect.Value, rtype.NumIn())

		for i := 0; i < rtype.NumIn(); i++ {

			if rtype.In(i).Implements(reflect.TypeOf([0]Parser{}).Elem()) && rtype.In(i).Kind() == reflect.Ptr {
				value := reflect.New(rtype.In(i).Elem())

				if err := value.Interface().(Parser).Parse(args[i].String()); err != nil {
					return handle(err)
				}

				converted[i] = value

			} else {

				switch rtype.In(i) {
				case reflect.TypeOf(time.Time{}):
					if args[i].Type() == js.TypeNumber {

						val := time.Millisecond * time.Duration(args[i].Float())
						converted[i] = reflect.ValueOf(time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).Add(val))

					} else {

						val := args[i].String()

						t, err := time.Parse(`2006-01-02`, val)
						if err != nil {
							t, err = time.Parse(`2006-01`, val)
							if err != nil {
								return handle(fmt.Errorf("bad request: %w", err))
							}
						}
						converted[i] = reflect.ValueOf(t)
					}

				default:
					switch rtype.In(i).Kind() {
					case reflect.String:
						converted[i] = reflect.ValueOf(args[i].String())
					default:
						return handle(errors.New("invalid arg-type " + rtype.In(i).String()))
					}
				}

			}
		}

		results := reflect.ValueOf(f).Call(converted)

		if len(results) == 0 {
			return nil
		}

		var result = results[0].Interface()

		type JSONEncodable interface {
			JSONEncoder() func(interface{}) ([]byte, error)
		}

		var buffer bytes.Buffer

		var encoder = json.NewEncoder(&buffer)

		switch v := result.(type) {
		case JSONEncodable:
			var encoder = v.JSONEncoder()

			b, err := encoder(result)
			if err != nil {
				return handle(fmt.Errorf("rpc function could not send return value: %w", err))
			}
			buffer.Write(b)

		case json.Marshaler:
			encoded, err := v.MarshalJSON()
			if err != nil {
				return handle(fmt.Errorf("rpc function could not send return value: %w", err))
			}
			buffer.Write(encoded)
		case qjs.AnyScript:
			var buffer bytes.Buffer
			ctx := qjs.NewCtx(&buffer)
			ctx(v.GetScript())
			ctx.Flush()

			return js.Global().Get("AsyncFunction").New(buffer.String()).Invoke()
		default:
			err := encoder.Encode(results[0].Interface())
			if err != nil {
				return handle(fmt.Errorf("invalid return-type %v: %v ", reflect.ValueOf(result).String(), err))
			}
		}

		return js.Global().Get("AsyncFunction").New(fmt.Sprintf(`return %v;`, buffer.String())).Invoke()
	}))
}
