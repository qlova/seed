package script

import qlova "github.com/qlova/script"

//Alert displays a message to the user.
func (q Ctx) Alert(message qlova.String) {
	q.js.Run("alert", message)
}

//Confirm displays a confirmation box to the user, returns a bool indicating true for 'ok' false for 'cancel'.
func (q Ctx) Confirm(message qlova.String) qlova.Bool {
	return q.js.Call("confirm", message).Bool()
}

//Prompt displays a prompt that requests a string from the user. This string is returned.
func (q Ctx) Prompt(message qlova.String) qlova.String {
	return q.js.Call("prompt", message).String()
}
