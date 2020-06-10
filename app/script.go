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
	return js.Func("launch").Run(location.Origin, js.NewString("installed"), js.NewNumber(800), js.NewNumber(600))
}
