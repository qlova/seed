package state

import (
	"fmt"
	"strconv"
)

//Sprintf returns a new String formatted with the provided values like fmt.Sprintf.
func Sprintf(format string, a ...Variable) String {
	var converted = make([]interface{}, len(a))
	var references = make([]Reference, len(a))
	for i, arg := range a {
		converted[i] = `"+(window.localStorage.getItem("` + arg.Ref() + `") || "")+"`
		references[i] = a[i].GetReference()
	}

	return String{
		Reference:  NewVariable(format),
		Expression: fmt.Sprintf(strconv.Quote(format), converted...),

		dependencies: references,
	}
}
