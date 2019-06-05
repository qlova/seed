package datebox

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"
import qlova "github.com/qlova/script"

import "github.com/qlova/seed/widgets/textbox"

type Widget struct {
	textbox.Widget
	fakebox textbox.Widget
}

func New() Widget {
	var TextBox = textbox.New()
	TextBox.SetAttributes("readonly")

	var FakeBox = textbox.AddTo(TextBox)
	FakeBox.SetAttributes("type='date'")
	FakeBox.Set("opacity", "0")
	FakeBox.Set("position", "absolute")
	FakeBox.Set("pointer-events", "none")
	FakeBox.OnChange(func(q seed.Script) {
		TextBox.Script(q).SetValue(FakeBox.Script(q).Value())
	})

	TextBox.OnChange(func(q seed.Script) {
		FakeBox.Script(q).SetValue(TextBox.Script(q).Value())
	})

	TextBox.OnClick(func(q seed.Script) {
		q.Javascript(`if (/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)) {`)
		FakeBox.Script(q).Focus()
		FakeBox.Script(q).Click()
		q.Javascript(`}`)
	})

	TextBox.OnReady(func(q seed.Script) {
		q.Javascript(`if (!(/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent))) {`)
		q.Javascript(TextBox.Script(q).Element() + `.readOnly = false;`)
		q.Javascript(TextBox.Script(q).Element() + `.type = "date";`)
		q.Javascript(`}`)
	})

	return Widget{TextBox, FakeBox}
}

func (widget Widget) SetRequired() {
	widget.fakebox.SetRequired()
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	parent.Root().Add(widget)
	return widget
}

func (widget Widget) Script(q seed.Script) Script {
	return Script{
		widget.Widget.Seed.Script(q),
		widget.fakebox.Seed.Script(q),
	}
}

type Script struct {
	script.Seed
	fakebox script.Seed
}

func (script Script) SetValue(value qlova.String) {
	script.Seed.SetValue(value)
	script.fakebox.SetValue(value)
}
