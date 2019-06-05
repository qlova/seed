package canvas

import qlova "github.com/qlova/script"
import "github.com/qlova/seed"
import "github.com/qlova/seed/gl"
import "github.com/qlova/seed/script"
import "github.com/qlova/seed/script/webgl"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()

	widget.SetTag("canvas")
	widget.SetSize(seed.Auto, seed.Auto)

	return Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}

//Return an OpenGL context for this canvas.
func (widget Widget) OpenGL() gl.Context {
	return gl.NewContext(widget.Seed)
}

type Script struct {
	script.Seed
}

func (w Widget) Script(q seed.Script) Script {
	return Script{w.Seed.Script(q)}
}

func (s Script) OpenGL() webgl.Context {
	return webgl.NewContext(s.Seed)
}

func (s Script) Width() qlova.Float {
	return s.Q.Value(s.Element() + ".scrollWidth").Float()
}

func (s Script) Height() qlova.Float {
	return s.Q.Value(s.Element() + ".scrollHeight").Float()
}
