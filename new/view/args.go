package view

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/js"
)

func valueAs(v js.AnyValue, T reflect.Type) reflect.Value {
	if T == reflect.TypeOf([0]client.Value{}).Elem() {
		return reflect.ValueOf(v)
	}

	var TypeName = strings.Replace(T.Name(), "Any", "", 1)
	if strings.Contains(TypeName, ".") {
		TypeName = strings.Split(TypeName, ".")[1]
	}

	if method, ok := T.MethodByName("Get" + TypeName); ok {
		Type := method.Type.Out(0)

		var result = reflect.New(Type).Elem()

		result.FieldByName("Value").Set(reflect.ValueOf(v.GetValue()))

		return result
	}

	return reflect.Zero(T)
}

//parseArgs returns the page arguments as a js.Object.
func parseArgs(view View, parent seed.Seed) (View, js.AnyObject) {
	if view == nil {
		return view, js.NewObject(nil)
	}

	var T = reflect.TypeOf(view)
	var V = reflect.ValueOf(view)

	var object = make(map[string]js.AnyValue, T.NumField())

	var NewView = reflect.New(T).Elem()

	for i := 0; i < T.NumField(); i++ {
		var Field = T.Field(i)
		var FieldValue = V.Field(i)
		if Field.Type.Implements(reflect.TypeOf((*js.AnyValue)(nil)).Elem()) {
			if intf := FieldValue.Interface(); intf != nil {

				var key = Field.Name

				object[key] = intf.(js.AnyValue)

				var value = js.NewValue(
					fmt.Sprintf(`seed.arg(`+html.Element(parent).String()+".CurrentView, %v)",
						strconv.Quote(key)))

				NewView.Field(i).Set(valueAs(value, Field.Type))
			}
		}
	}

	return NewView.Interface().(View), js.NewObject(object)
}
