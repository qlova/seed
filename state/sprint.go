package state

import (
	"fmt"
	"strconv"
)

//Sprintf returns a new String formatted with the provided values like fmt.Sprintf.
func Sprintf(format string, a ...AnyValue) String {
	var converted = make([]interface{}, len(a))
	var references = make([]Value, len(a))
	for i, arg := range a {
		converted[i] = `"+` + arg.value().getter() + `+"`
		references[i] = a[i].value()
	}

	var v = newValue(format)
	v.dependencies = &references
	v.raw = fmt.Sprintf(strconv.Quote(format), converted...)
	return String{v}
}
