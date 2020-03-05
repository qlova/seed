package app

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/app/manifest"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/script"
)

//App is a webapp generator.
type App struct {
	seed.Seed

	manifest manifest.Manifest

	document html.Document

	name string

	page page.Seed

	description string
}

//New returns a new App.
func New(name string, options ...seed.Option) App {
	var document = html.New()

	var app = App{
		Seed:     document.Body,
		document: document,
		name:     name,
		manifest: manifest.New(),
	}

	app.manifest.SetName(name)

	for _, o := range options {
		o.AddTo(app)
	}

	return app
}

//SetPage sets the default page for this app.
func SetPage(page page.Page) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		if app, ok := s.(App); ok {
			script.OnReady(func(q script.Ctx) {
				//page.Ctx(q).SetStartingPage()
			}).AddTo(app)
		}
	}, nil, nil)
}
