package filepicker

import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

func New(types ...string) Widget {
	widget := seed.New()
	widget.SetTag("Input")

	if len(types) > 0 {
		widget.SetAttributes(`type="file" accept="`+types[0]+`"`)
	} else {
		widget.SetAttributes(`type="file" accept="*"`)
	}

	return  Widget{widget}
}

func AddTo(parent seed.Interface, types ...string) Widget {
	var widget = New(types...)
	parent.Root().Add(widget)
	return widget
}
