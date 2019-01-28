package button

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New(text ...string) Widget {
	widget := seed.New()
	
	widget.SetTag("button")
	
	widget.SetSize(seed.Auto, seed.Auto)
	
	if len(text) > 0 {
		widget.SetText(text[0])
	}

	return Widget{widget}
}

func AddTo(parent seed.Interface, text ...string) Widget {
	var widget = New(text...)
	parent.Root().Add(widget)
	return widget
}
