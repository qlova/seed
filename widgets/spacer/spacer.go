package spacer

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New(amount ...complex128) Widget {
	widget := seed.New()

	if len(amount) > 0 {
		widget.SetSize(amount[0], amount[0])
	}

	widget.SetUnshrinkable()
	
	return  Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, amount ...complex128) Widget {
	var widget = New(amount...)
	parent.Root().Add(widget)
	return widget
} 
