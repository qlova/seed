//Package echo provides powerful ways to modify complex objects within a script context.
package echo

import (
	"reflect"
	"strconv"

	"qlova.org/seed/script"
)

//Ctx allows script Values to be assigned to fields inside of t ChangesTo function.
type Ctx func(script.AnyValue) interface{}

func (c Ctx) String(v script.AnyValue) string {
	return c(v).(string)
}

//ChangesTo takes a function of the form func(Ctx, *T)
//Any changes made to *T using script values will be replicated on the provided value.
func ChangesTo(v script.Value, f interface{}) script.Script {
	return func(q script.Ctx) {
		var FunctionType = reflect.TypeOf(f)
		var FunctionValue = reflect.ValueOf(f)

		var Values []script.Value

		if FunctionType.NumIn() == 2 {
			//Prepare the echo-location equipment.
			var ctx = func(any script.AnyValue) interface{} {
				value := any.GetValue()
				Values = append(Values, value)
				switch any.(type) {
				case script.String:
					return strconv.Itoa(len(Values))
				default:
					panic("unexpected type in call to echo.Ctx: " + reflect.TypeOf(any).String())
				}
			}
			var echo = reflect.New(FunctionType.In(1).Elem())

			//Send out the echo
			FunctionValue.Call([]reflect.Value{reflect.ValueOf(ctx), echo})

			//Inspect the echo.
			var equipment = echo.Elem()
			var T = equipment.Type()

			switch T.Kind() {
			case reflect.Struct:
				for i := 0; i < T.NumField(); i++ {
					var index int = -1

					var field = equipment.Field(i)

					switch field.Type().Kind() {
					case reflect.String:
						index, _ = strconv.Atoi(field.String())
					}

					if index > -1 {
						q(v.Set(T.Field(i).Name, Values[index-1]))
					}
				}
			default:
				panic("unexpected type in call to echo.ChangesTo: " + T.String())
			}

		}
	}
}
