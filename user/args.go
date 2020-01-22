package user

import (
	"encoding/json"
	"strings"
)

//Arg is a user provided argument.
type Arg struct {
	string
}

//Arg returns the argument with the given name.
func (u Ctx) Arg(name string) Arg {
	return Arg{u.r.FormValue(name)}
}

//String returns the argument as a string.
func (arg Arg) String() string {
	return arg.string
}

//Strings returns the argument as a slice of strings.
func (arg Arg) Strings() (returns []string, err error) {
	err = json.NewDecoder(strings.NewReader(arg.string)).Decode(&returns)
	return
}

//InterfaceMap returns the argument as a map of interface values.
func (arg Arg) InterfaceMap() (returns map[string]interface{}, err error) {
	err = json.NewDecoder(strings.NewReader(arg.string)).Decode(&returns)
	return
}

//Decode decodes the value into the given argument.
func (arg Arg) Decode(i interface{}) error {
	return json.NewDecoder(strings.NewReader(arg.string)).Decode(i)
}
