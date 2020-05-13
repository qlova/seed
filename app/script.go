package app

import (
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/js/location"

	"github.com/qlova/seed/script"
)

//Reset resets the app and clears any local storage.
func Reset() script.Script {
	return func(q script.Ctx) {
		q(`window.sessionStorage.clear(); window.localStorage.clear(); window.location = "/";`)
	}
}

//Launch launches the current app as new window in an installed state.
func Launch() script.Script {
	var window = js.Global().Get("window")
	return window.Run("open", location.Origin, js.NewString("installed"), js.NewString("height=800,width=600"))
}
