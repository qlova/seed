package user

import (
	"encoding/json"
	"log"
	"strings"
)

//Argument is a user provided argument.
type Argument struct {
	string
}

//Args returns the argument with the given name.
func (user User) Args(name string) Argument {
	return Argument{user.Request.FormValue(name)}
}

func (arg Argument) String() string {
	return arg.string
}

//Strings returns the argument as a slice of strings.
func (arg Argument) Strings() (returns []string) {
	err := json.NewDecoder(strings.NewReader(arg.string)).Decode(&returns)
	if err != nil {
		log.Println(err)
	}
	return
}

//Map returns the argument as a map of interface values.
func (arg Argument) Map() (returns map[string]interface{}) {
	err := json.NewDecoder(strings.NewReader(arg.string)).Decode(&returns)
	if err != nil {
		log.Println(err)
	}
	return
}
