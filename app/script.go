package app

import (
	"qlova.org/seed/client"
	"qlova.org/seed/js"
	"qlova.org/seed/js/console"
	"qlova.org/seed/js/location"

	"qlova.org/seed/script"
)

//Reset resets the app and clears any local storage.
func Reset() script.Script {
	return func(q script.Ctx) {
		q(`window.sessionStorage.clear(); window.localStorage.clear();  window.location = "/";`)
	}
}

//Update updates the app.
func Update() script.Script {
	return script.New(console.Log(client.NewString("update")),
		func(q script.Ctx) {
			q(`document.cookie = "version=; max-age=-1;";`)
			q(`if (window.ServiceWorker_Registration) await ServiceWorker_Registration.update();`)
		},
		Restart(),
	)
}

//Restart restarts the app.
func Restart() script.Script {
	return func(q script.Ctx) {
		q(`window.location = "/";`)
	}
}

//Launch launches the current app as new window in an installed state.
func Launch() client.Script {
	return js.Func("launch").Run(location.Origin, js.NewString("installed"), js.NewNumber(800), js.NewNumber(600))
}

//Install installs the app as a PWA installed application on supported browsers.
func Install() client.Script {
	return client.NewScript(
		Installable.Set(false),
		js.Func("install").Run(),
	)
}
