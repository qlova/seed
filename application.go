package seed

import "github.com/qlova/seed/manifest"

type Application struct {
	Seed
	manifest.Manifest
}

func NewApp() Application {
	return Application {
		Seed: New(),
		Manifest: manifest.New(),
	}
}

//TODO random port, can be set with enviromental variables.
func (app Application) Launch(listen ...string) error {
	Launcher{Application: app}.Launch(listen...)
	return nil
}
