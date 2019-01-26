package spacer

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New(amount ...float64) Widget {
	widget := seed.New()

	if len(s) > 0 {
		widget.SetSize(s[0], s[0])
	}

	return  Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, amount ...float64) Widget {
	var widget = New(amount...)
	parent.Root().Add(widget)
	return widget
} 
