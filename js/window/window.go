package window

import (
	"github.com/qlova/seed/js"
)

//Alert creates a popup window with the provided text message.
func Alert(msg js.AnyString) js.Script {
	return func(q js.Ctx) {
		q.Run(`window.alert`, msg)
	}
}

//Confirm creates a popup window with the provided text message asking the user to confirm.
func Confirm(msg js.AnyString) js.Bool {
	return js.Bool{Value: js.Call(`window.confirm`, msg)}
}

//Close attempts to close the window.
func Close() js.Script {
	return func(q js.Ctx) {
		q.Run(`window.close()`)
	}
}

//Prompt creates a popup window with the provided text message asking the user to input text.
func Prompt(msg js.AnyString) js.String {
	return js.String{Value: js.Call(`window.prompt`, msg)}
}
