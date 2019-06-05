package document

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New(path ...string) Widget {
	widget := seed.New()

	widget.SetTag("embed")
	if len(path) > 0 {
		widget.SetAttributes("src='" + path[0] + "'")
		seed.NewAsset(path[0]).AddTo(widget)
	}

	return Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, path ...string) Widget {
	var widget = New(path...)
	parent.Root().Add(widget)
	return widget
}
