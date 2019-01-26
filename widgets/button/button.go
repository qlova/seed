package button

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()
	
	widget.SetTag("button")
	
	widget.SetSize(seed.Auto, seed.Auto)

	return Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}
