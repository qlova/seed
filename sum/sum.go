package sum

import (
	"fmt"
	"reflect"

	"qlova.org/seed/js"

	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
)

//String will be treated as a string.
type String interface{}

//ToString converts the value into a (string, client.String) combo
func ToString(s String) (string, client.String) {
	switch v := s.(type) {
	case string:
		return v, nil

	case *clientside.String:
		return v.Value, v

	case clientside.String:
		panic("sum.ToString: expected *clientside.String but got clientside.String")

	case client.String:
		return "", v

	case client.Value:
		return "", js.String{Value: js.Func("String").Call(v)}

	default:
		return fmt.Sprint(s), nil
	}
}

//Float64 will be treated as a float64.
type Float64 interface{}

//ToFloat64 converts the value into a (string, client.String) combo
func ToFloat64(f Float64) (float64, client.Float) {
	switch v := f.(type) {
	case float64:
		return v, nil

	case *clientside.Float64:
		return v.Value, v

	case clientside.Float64:
		panic("sum.ToFloat64: expected *clientside.Float64 but got clientside.Float64")

	case client.Float:
		return 0, v

	case client.Value:
		return 0, js.Number{Value: js.Func("Number").Call(v)}

	default:
		panic("sum.ToFloat64: expected floaty value but got " + reflect.TypeOf(f).String())
	}
}
