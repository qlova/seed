package emailbox

import "github.com/qlova/seed"
import "github.com/qlova/seed/widgets/textbox"

type Widget struct {
	textbox.Widget
}

func New() Widget {
	widget := textbox.New()
	widget.SetAttributes(`type="email"`)

	return  Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}