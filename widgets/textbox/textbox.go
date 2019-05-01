package textbox

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()
	widget.SetTag("input")
	
	widget.SetSize(seed.Auto, seed.Auto)

	widget.Align(0)

	var save = script.NewString()
	widget.OnChange(func(q seed.Script) {
		q.Set(save, widget.Script(q).Value())
	})
	widget.OnReady(func(q seed.Script) {
		widget.Script(q).SetValue(save.Script(q))
	})

	return  Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}

func (widget Widget) SetRequired() {
	widget.SetAttributes(widget.Attributes()+" required")
}


type Script struct {
	script.Seed
}

func (w Widget) Script(q script.Script) Script {
	return Script{w.Seed.Script(q)}
}