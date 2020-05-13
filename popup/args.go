package popup

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
func parseArgs(popup Popup) (Popup, js.AnyObject) {
	if popup == nil {
		return popup, js.NewObject(nil)
	}

	var T = reflect.TypeOf(popup)
	var V = reflect.ValueOf(popup)

	var object = make(map[string]js.AnyValue, T.NumField())

	var NewPopup = reflect.New(T).Elem()

	for i := 0; i < T.NumField(); i++ {
		var Field = T.Field(i)
		var FieldValue = V.Field(i)
		if Field.Type.Implements(reflect.TypeOf((*js.AnyValue)(nil)).Elem()) {
			if intf := FieldValue.Interface(); intf != nil {

				var key = Field.Name

				object[key] = intf.(js.AnyValue)

				var value = js.NewValue(
					fmt.Sprintf("seed.get(%v).args[%v]", strconv.Quote(ID(popup)),
						strconv.Quote(key)))

				NewPopup.Field(i).Set(valueAs(value, Field.Type))
			}
		}
	}

	return NewPopup.Interface().(Popup), js.NewObject(object)
}
