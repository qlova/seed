package location

import (
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

//Origin is the origin part of the current location.
var Origin = js.String{Value: js.NewValue(`window.location.origin`)}

func Replace(url js.AnyString) script.Script {
	return js.Global().Run("location.replace", url)
}

func Reload() script.Script {
	return js.Global().Run("location.reload")
}
