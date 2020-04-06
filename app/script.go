package app

import "github.com/qlova/seed/script"

//Reset resets the app and clears any local storage.
func Reset() script.Script {
	return func(q script.Ctx) {
		q(`window.sessionStorage.clear(); window.localStorage.clear(); window.location = "/";`)
	}
}
