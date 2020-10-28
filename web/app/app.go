package app

import (
	"image/color"

	"qlova.org/seed"
	"qlova.org/seed/web/app/manifest"
	"qlova.org/seed/web/app/service"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/web/css"
	"qlova.org/seed/web/html"
	"qlova.org/seed/web/js"
	"qlova.org/seed/new/page"
)

type App struct {
	seed.Seed

	port string
}

//App is a webapp generator.
type app struct {
	manifest manifest.Manifest
	worker   *service.Worker

	document html.Document

	name string

	page, loadingPage page.Page

	description string

	color color.Color
}

//Installable is true when the app can be installed (that is when the OS has granted the app a beforeinstallprompt event).
var Installable = &clientside.Bool{
	Name: "app.installable",
}

//Standalone is true when the app is running from an installed instance.
var Standalone = &clientside.Bool{
	Name:  "app.standalone",
	Value: js.NewValue(`(document.fullscreen || (window.matchMedia('(display-mode: standalone)').matches) || (window.navigator.standalone) || window.name == 'installed')`),
}

//New returns a new App.
func New(name string, options ...seed.Option) App {
	var document = html.New()

	var SeedCount = 0
	for i := range options {
		if _, ok := options[i].(seed.Seed); ok {
			SeedCount++
		}
	}

	var app = app{
		document: document,
		name:     name,
		manifest: manifest.New(),
		worker:   service.NewWorker(),
	}

	document.Body.With(css.Set("display", "flex"))
	document.Body.With(css.Set("flex-direction", "column"))

	app.manifest.SetName(name)

	app.manifest.Icons = append(app.manifest.Icons, manifest.Icon{
		Source: "/Qlovaseed.png",
		Sizes:  "512x512",
	})

	document.Body.Save(app)

	for _, o := range options {
		o.AddTo(document.Body)
	}

	document.Body.Load(&app)

	document.Save(app)

	return App{document.Seed, ""}
}
