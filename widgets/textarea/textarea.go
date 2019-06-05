package textarea

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()
	widget.SetTag("textarea")
	widget.SetAttributes("data-gramm_editor=false")

	return Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}

func (widget Widget) SetRequired() {
	widget.SetAttributes(widget.Attributes() + " required")
}
