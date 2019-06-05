package slider

import "strconv"
import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()
	widget.SetTag("input")
	widget.SetAttributes("type='range'")

	widget.SetSize(seed.Auto, seed.Auto)

	widget.Align(0)

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

func (widget Widget) SetMax(max int) {
	widget.SetAttributes(widget.Attributes() + " max='" + strconv.Itoa(max) + "'")
}

type Script struct {
	script.Seed
}

func (w Widget) Script(q script.Script) Script {
	return Script{w.Seed.Script(q)}
}
