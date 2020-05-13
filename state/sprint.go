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
		converted[i] = fmt.Sprintf(`"+%v+"`, arg.value().getter())
		references[i] = a[i].value()
	}

	var v = newValue(format)
	v.dependencies = &references
	v.raw = fmt.Sprintf(strconv.Quote(format), converted...)
	return String{v}
}

//Or returns a Bool that is true when either are true.
func (s State) Or(or State) State {
	return State{Bool: s.Bool.Or(or.Bool)}
}

//And returns a Bool that is true when both are true.
func (s State) And(or State) State {
	return State{Bool: s.Bool.And(or.Bool)}
}
