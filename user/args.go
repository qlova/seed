package user

import (
	"encoding/json"
	"strconv"
	"strings"
)

//Arg is a user provided argument.
type Arg struct {
	u Ctx
	i string
}

//Arg returns the argument with the given name.
func (u Ctx) Arg(name string) Arg {
	return Arg{u, name}
}

//String returns the argument as a string.
func (arg Arg) String() string {
	n, err := strconv.Atoi(arg.i)
	if err == nil {
		return arg.u.r.FormValue(string('a' + rune(n)))
	}
	return arg.u.r.FormValue(arg.i)
}

//Strings returns the argument as a slice of strings.
func (arg Arg) Strings() (returns []string, err error) {
	err = json.NewDecoder(strings.NewReader(arg.String())).Decode(&returns)
	return
}

//InterfaceMap returns the argument as a map of interface values.
func (arg Arg) InterfaceMap() (returns map[string]interface{}, err error) {
	err = json.NewDecoder(strings.NewReader(arg.String())).Decode(&returns)
	return
}

//Decode decodes the value into the given argument.
func (arg Arg) Decode(i interface{}) error {
	return json.NewDecoder(strings.NewReader(arg.String())).Decode(i)
}
