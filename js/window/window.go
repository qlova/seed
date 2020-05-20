package window

import (
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

//Alert creates a popup window with the provided text message.
func Alert(msg js.AnyString) js.Script {
	return js.Global().Run(`alert`, msg)
}

//Confirm creates a popup window with the provided text message asking the user to confirm.
func Confirm(msg js.AnyString) js.Bool {
	return js.Bool{js.Global().Call(`confirm`, msg)}
}

//Close attempts to close the window.
func Close() js.Script {
	return js.Global().Run(`close`)
}

//Prompt creates a popup window with the provided text message asking the user to input text.
func Prompt(msg js.AnyString) js.String {
	return js.String{js.Global().Call(`prompt`, msg)}
}

func SetTimeout(do js.Script, timeout js.Number) script.Script {
	return js.Global().Run("setTimeout", js.NewFunction(do), timeout)
}

//ResizeTo dynamically resizes the window.
func ResizeTo(width, height js.Number) js.Script {
	return js.Global().Run(`resizeTo`, width, height)
}
