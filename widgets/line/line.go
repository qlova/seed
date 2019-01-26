package line

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()

	widget.SetTag("hr")
	
	widget.SetSize(seed.Auto, seed.Auto)

	widget.Set("border-style", "solid")

	return  Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
} 
