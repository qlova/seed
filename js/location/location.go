package location

import (
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

//Origin is the origin part of the current location.
var Origin = js.String{Value: js.NewValue(`window.location.origin`)}

func Replace(url js.AnyString) script.Script {
	return js.Global().Run("location.replace", url)
}

func Reload() script.Script {
	return js.Global().Run("location.reload")
}
