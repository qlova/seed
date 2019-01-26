package seed

import "github.com/qlova/seed/manifest"

type App struct {
	Seed
	manifest.Manifest
}

//Create a new application, accepts title and content arguments.
func NewApp(args ...string) App {
	var app = App {
		Seed: New(),
		Manifest: manifest.New(),
	}

	if len(args) > 0 {
		app.SetName(args[0])
	}

	if len(args) > 1 {
		app.SetContent(args[1])
	}

	return app
}

//TODO random port, can be set with enviromental variables.
func (app App) Launch(listen ...string) error {
	launcher{App: app}.Launch(listen...)
	return nil
}
