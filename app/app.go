package app

import (
	"image/color"

	"qlova.org/seed"
	"qlova.org/seed/app/manifest"
	"qlova.org/seed/app/service"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
	"qlova.org/seed/page"
)

type App struct {
	seed.Seed

	port string
}

//App is a webapp generator.
type app struct {
	seed.Data

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

		if _, ok := options[i].(page.Seed); ok {
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

	document.Body.Write(app)

	for _, o := range options {
		o.AddTo(document.Body)
	}

	document.Body.Read(&app)

	document.Write(app)

	return App{document, ""}
}
