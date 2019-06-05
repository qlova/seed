package row

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()

	widget.Stylable.Set("display", "flex")
	widget.Stylable.Set("flex-direction", "row")
	widget.Stylable.Set("flex-shrink", "1")

	return Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}
