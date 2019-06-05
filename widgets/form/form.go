package form

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"
import . "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()
	widget.SetTag("form")
	widget.SetAttributes(`onsubmit="return false;"`)

	widget.Stylable.Set("display", "flex")
	widget.Stylable.Set("flex-direction", "column")

	return Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}

type Script struct {
	script.Seed
}

func (w Widget) Script(q script.Script) Script {
	return Script{w.Seed.Script(q)}
}

func (widget Script) Invalid() Bool {
	return widget.Q.BoolFromLanguageType(Javascript.Bit{
		Expression: language.Statement("!" + widget.Element() + ".reportValidity()"),
	})
}
