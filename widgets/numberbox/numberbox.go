package numberbox

import "github.com/qlova/seed"

import "github.com/qlova/seed/widgets/textbox"

type Widget struct {
	textbox.Widget
}

func New() Widget {
	var TextBox = textbox.New()
	TextBox.SetAttributes("type='number'")

	return Widget{TextBox}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}
