package clientargs

import (
	"reflect"
	"strings"

	"github.com/qlova/seed/client"
	"github.com/qlova/seed/js"
)

//PassGoValues can be embedded inside of a clientside-arguments struct to enable the passthrough of Go arguments.
//these arguments must be decoded with this type's methods after being parsed.
type PassGoValues struct {
	arguments client.Object
	ints      int
}

func (p PassGoValues) getPassGoValues() PassGoValues {
	return p
}

//Is returns a client reference to 'true' if the given popup field is equal to the provided value.
func (p PassGoValues) Is(field interface{}, value interface{}) client.Bool {
	var T = reflect.TypeOf(field)
	var V = reflect.ValueOf(field)

	switch T.Kind() {
	case reflect.Int:
		return p.arguments.GetObject().Value.Get("types").Get("int").Index(int(V.Int()) - 1).Equals(js.ValueOf(value))
	}

	return client.NewBool(false)
}

func valueAs(v js.AnyValue, T reflect.Type) reflect.Value {
	var TypeName = strings.TrimPrefix(T.Name(), "Any")
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

//Parse parses the arguments of a clientside-arguments struct and returns the converted client-side arguments struct as an interface and the client representation of it.
func Parse(structure interface{}, arguments client.Object) (interface{}, client.Object) {

	type passingGoValues interface {
		getPassGoValues() PassGoValues
	}

	getter, passing := structure.(passingGoValues)
	var pgv PassGoValues
	if passing {
		pgv = getter.getPassGoValues()
		pgv.arguments = arguments
	}

	var T = reflect.TypeOf(structure)
	var V = reflect.ValueOf(structure)

	var object = make(map[string]js.AnyValue, T.NumField())
	var TypeTable = js.NewObject{}
	if object["types"] == nil {
		object["types"] = TypeTable
	}

	var Converted = reflect.New(T).Elem()

	for i := 0; i < T.NumField(); i++ {
		var Field = T.Field(i)
		var FieldValue = V.Field(i)
		if Field.Type.Implements(reflect.TypeOf((*js.AnyValue)(nil)).Elem()) {

			var key = Field.Name

			if intf := FieldValue.Interface(); intf != nil {
				object[key] = intf.(js.AnyValue)
			} else {
				object[key] = js.Null()
			}

			var value = arguments.GetObject().Get(client.NewString(key))

			Converted.Field(i).Set(valueAs(value, Field.Type))

		} else if passing {

			if _, ok := FieldValue.Interface().(passingGoValues); ok {
				val := FieldValue
				cval := Converted.Field(i)
				for {
					if val.Type() == reflect.TypeOf(PassGoValues{}) {
						cval.Set(reflect.ValueOf(pgv))
						break
					}
					val = val.Field(0)
					cval = cval.Field(0)
				}
				continue
			}

			switch Field.Type.Kind() {
			case reflect.Int:
				pgv.ints++
				if TypeTable["int"] == nil {
					TypeTable["int"] = js.NewArray{}
				}
				ints := TypeTable["int"].(js.NewArray)
				TypeTable["int"] = append(ints, js.NewNumber(float64(FieldValue.Int())))

				Converted.Field(i).SetInt(int64(pgv.ints))
			default:
				panic("clientargs.Parse: unsupported type: " + Field.Type.String())
			}

		}

	}

	return Converted.Interface(), js.NewObject(object)
}
