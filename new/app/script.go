package app

import (
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
	"qlova.org/seed/use/js/console"
	"qlova.org/seed/use/js/location"
)

//Reset resets the app and clears any local storage.
func Reset() client.Script {
	return js.Script(func(q js.Ctx) {
		q(`window.sessionStorage.clear(); window.localStorage.clear();  window.location = "/";`)
	})
}

//Update updates the app.
func Update() client.Script {
	return client.NewScript(console.Log(client.NewString("update")),
		js.Script(func(q js.Ctx) {
			q(`document.cookie = "version=; max-age=-1;";`)
			q(`if (window.ServiceWorker_Registration) await ServiceWorker_Registration.update();`)
		}),
		Restart(),
	)
}

//Restart restarts the app.
func Restart() client.Script {
	return js.Script(func(q js.Ctx) {
		q(`window.location.reload();`)
	})
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
