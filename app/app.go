package app

import (
	"image/color"

	"github.com/qlova/seed"
	"github.com/qlova/seed/app/manifest"
	"github.com/qlova/seed/app/service"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/state"
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

//Installed is true when the app is running from an installed instance.
var Installed = state.State{
	Bool: state.Bool{
		Value: state.Raw("(window.matchMedia('(display-mode: standalone)').matches) || (window.navigator.standalone)"),
	},
}

//New returns a new App.
func New(name string, options ...seed.Option) App {
	var document = html.New()

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
