package checkbox

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()

	widget.SetTag("input")
	widget.SetAttributes("type='checkbox'")

	return Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}
