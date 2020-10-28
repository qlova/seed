package location

import (
	"qlova.org/seed/web/js"
)

//Origin is the origin part of the current location.
var Origin = js.String{Value: js.NewValue(`window.location.origin`)}

func Replace(url js.AnyString) js.Script {
	return js.Global().Run("location.replace", url)
}

func Reload() js.Script {
	return js.Global().Run("location.reload")
}
