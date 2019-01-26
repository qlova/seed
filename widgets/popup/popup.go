package popup

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()
	
	widget.Set("position", "fixed")
	widget.Set("left", "50%")
	widget.Set("top", "50%")
	widget.Set("transform", "translate(-50%, -50%)")
	widget.Set("box-shadow", "3px 4px 20px black")

	widget.SetSize(seed.Auto, seed.Auto)
	widget.SetHidden()

	return Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}
