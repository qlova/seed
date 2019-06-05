package expander

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New(ratio ...float64) Widget {
	widget := seed.New()

	if len(ratio) > 0 {
		widget.SetExpand(ratio[0])
	} else {
		widget.SetExpand(1)
	}

	return Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, ratio ...float64) Widget {
	var widget = New(ratio...)
	parent.Root().Add(widget)
	return widget
}
