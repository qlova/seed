package app

import (
	"image/color"

	"qlova.org/seed"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/new/app/manifest"
	"qlova.org/seed/new/app/service"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set/center"
	"qlova.org/seed/use/css"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/js"
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

	name, description, pkg string

	hashes []string

	page, loadingPage page.Page

	color color.Color

	head []seed.Option
}

//Installable is true when the app can be installed (that is when the OS has granted the app a beforeinstallprompt event).
var Installable = &clientside.Bool{
	MemoryAddress: clientside.MemoryAddress{
		Name: "app.installable",
	},
}

//Standalone is true when the app is running from an installed instance.
var Standalone = &clientside.Bool{
	MemoryAddress: clientside.MemoryAddress{
		Name: "app.standalone",
	},
	Value: js.NewValue(`(document.fullscreen || (window.matchMedia('(display-mode: standalone)').matches) || (window.navigator.standalone) || window.name == 'installed')`),
}

//New returns a new App.
func New(name string, options ...seed.Option) App {
	var document = html.New()

	//Make a little 'Hello World' app using the title.
	if len(options) == 0 {
		options = []seed.Option{
			center.This(
				text.New(text.SetString(name)),
			),
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
