package page

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/qlova/seed/js"
)

func valueAs(v js.AnyValue, T reflect.Type) reflect.Value {
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
func parseArgs(page Page) (Page, js.Object, js.String) {
	var url = js.NewString("")

	if page == nil {
		return page, js.NewObject(nil), url
	}

	var T = reflect.TypeOf(page)
	var V = reflect.ValueOf(page)

	var object = make(map[string]js.AnyValue, T.NumField())

	var NewPage = reflect.New(T).Elem()

	for i := 0; i < T.NumField(); i++ {
		var Field = T.Field(i)
		var FieldValue = V.Field(i)

		if Field.Type.Implements(reflect.TypeOf((*js.AnyValue)(nil)).Elem()) {
			if intf := FieldValue.Interface(); intf != nil {

				var key string

				//Simple base lookup.
				if tag, ok := Field.Tag.Lookup("url"); ok {
					key = tag
					switch tag {
					case `1`:
						url = js.String{Value: js.Call(`"/"+encodeURIComponent`, intf.(js.AnyValue))}
					default:
						panic("not implimented")
					}
				} else {
					key = Field.Name
				}

				object[key] = intf.(js.AnyValue)

				var value = js.NewValue(
					fmt.Sprintf("seed.CurrentPage.args[%v]",
						strconv.Quote(key)))

				NewPage.Field(i).Set(valueAs(value, Field.Type))
			}
		}
	}

	return NewPage.Interface().(Page), js.NewObject(object), url
}
