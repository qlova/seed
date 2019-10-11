package user

import "bytes"

//Update describes a JSON update that gets sent to the user, it defines a change in state that will be applied to the user.
type Update struct {

	//a mapping of document paths to values,
	// eg. #gg.style.display => hidden;
	Document map[string]string

	//a mapping of local-storage updates.
	LocalStorage map[string]string

	//A list of js statements that will be evaluated.
	Evaluations string

	//Script to evaluate.
	script *bytes.Buffer

	//A map of data to values that will be updated.
	Data map[string]string
}

func (update Update) Script() *bytes.Buffer {
	return update.script
}
