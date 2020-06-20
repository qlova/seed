package page

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"qlova.org/seed/js"
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
func parseArgs(page Page) (Page, js.AnyObject, js.String) {
	var url string = `""`
	var queries []string

	var T = reflect.TypeOf(page)
	var V = reflect.ValueOf(page)

	if page == nil || T.Kind() != reflect.Struct {
		return page, js.NewObject(nil), js.NewString(url)
	}

	var object = make(map[string]js.AnyValue, T.NumField())

	var NewPage = reflect.New(T).Elem()

	for i := 0; i < T.NumField(); i++ {
		var Field = T.Field(i)
		var FieldValue = V.Field(i)

		if Field.Type.Implements(reflect.TypeOf((*js.AnyValue)(nil)).Elem()) {
			intf := FieldValue.Interface()

			var key string

			//Simple base lookup.
			if tag, ok := Field.Tag.Lookup("url"); ok {
				key = tag
				if intf != nil {
					switch tag {
					case `1`:
						if url == `""` {
							url = `"/"`
						}
						url += js.Call(js.Function{js.NewValue(`+encodeURIComponent`)}, intf.(js.AnyValue)).String()
					default:
						queries = append(queries, js.Call(js.Function{js.NewValue(`"` + tag + `="+encodeURIComponent`)}, intf.(js.AnyValue)).String())
					}
				}
			} else {
				key = Field.Name
			}

			if intf != nil {
				object[key] = intf.(js.AnyValue)
			}

			var value = js.NewValue(
				fmt.Sprintf("seed.arg(%v, %v)", strconv.Quote(ID(page)),
					strconv.Quote(key)))

			NewPage.Field(i).Set(valueAs(value, Field.Type))
		}
	}

	//Build query.
	if len(queries) > 0 {
		url += `+"?"`
	}
	for i, q := range queries {
		url += `+` + q
		if i < len(queries)-1 {
			url += `+"&"`
		}
	}

	return NewPage.Interface().(Page), js.NewObject(object), js.String{Value: js.NewValue(url)}
}
