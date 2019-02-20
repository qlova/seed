package link

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"
import . "github.com/qlova/script"

type Widget struct {
	seed.Seed
}

func New(url ...string) Widget {
	widget := seed.New()
	widget.SetTag("a")
	
	if len(url) > 0 {
		widget.SetAttributes("href='"+url[0]+"'")
	} else {
		widget.SetAttributes("href='#'")
	}

	return Widget{widget}
}

func AddTo(parent seed.Interface, url ...string) Widget {
	var widget = New(url...)
	parent.Root().Add(widget)
	return widget
}

type Script struct {
	script.Seed
}

func (w Widget) Script(q script.Script) Script {
	return Script{w.Seed.Script(q)}
}

func (widget Script) SetTarget(target String) {
	widget.Q.Javascript(widget.Element()+`.href = `+string(target.LanguageType().Raw())+";")
}
