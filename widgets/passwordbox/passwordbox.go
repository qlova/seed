package passwordbox

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()
	widget.SetTag("input")
	widget.SetAttributes(`type="password"`)

	widget.SetSize(seed.Auto, seed.Auto)

	widget.Align(0)

	return Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}
