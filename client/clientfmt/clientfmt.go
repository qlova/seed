package clientfmt

import (
	"fmt"
	"strconv"

	"qlova.org/seed/client"
	"qlova.org/seed/js"
)

//String is a client formatted string.
type String struct {
	client.String
	components []client.Value
}

func NewString(s client.String, components ...client.Value) String {
	return String{s, client.FlattenComponents(components)}
}

//Components implements clientside.Compound
func (s String) Components() []client.Value {
	return s.components
}

//Concat returns a+b
func Concat(a, b js.AnyString) String {
	return NewString(js.String{js.NewValue("(%v + %v)", a, b)}, a, b)
}

//Sprintf replaces the "%v" values in the fmt string with the given client values.
func Sprintf(format string, args ...interface{}) String {

	var components []client.Value

	var nargs = make([]interface{}, len(args))
	for i := range args {
		val, ok := args[i].(client.Value)
		if ok {
			nargs[i] = val.GetValue().String()
			components = append(components, val)
			continue
		}
		nargs[i] = args[i]
	}

	var converted = make([]interface{}, len(nargs))
	for i, arg := range nargs {
		converted[i] = fmt.Sprintf(`"+%v+"`, arg)
	}

	return String{
		String:     js.String{Value: js.NewValue(fmt.Sprintf("("+strconv.Quote(format)+")", converted...))},
		components: components,
	}
}
