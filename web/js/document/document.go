package document

import (
	"qlova.org/seed/web/js"
)

//ExitFullscreen asks the user agent to place the specified element (and, by extension, its descendants) into full-screen mode, removing all of the browser's UI elements as well as all other applications from the screen.
func ExitFullscreen() js.Script {
	return js.Global().Get("document").Run("exitFullscreen")
}
