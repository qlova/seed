package user

import (
	"bytes"
	"encoding/json"
	"log"
)

//Update describes a JSON update that gets sent to the user, it defines a change in state that will be applied to the user.
type Update struct {
	//Response is the return value.
	Response *string

	//a mapping of document paths to values,
	// eg. #gg.style.display => hidden;
	Document map[string]string

	//a mapping of local-storage updates.
	LocalStorage map[string]string

	//js script that will be evaluated.
	Evaluation string

	//Script to evaluate.
	script *bytes.Buffer

	//A map of data to values that will be updated.
	Data map[string]string
}

//Return a value to the user.
func (update Update) Return(value interface{}) {

	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(value)
	if err != nil {
		log.Println("could not encode: ", err)
	}
	*update.Response = buffer.String()
}

func (update Update) Script() *bytes.Buffer {
	return update.script
}
