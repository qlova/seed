package checkbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"

	"qlova.org/seed/new/html/input"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(
		attr.Set("type", "checkbox"),
		seed.Options(options),
	)
}

//Update updates the given variable whenever the checkbox value is modified.
func Update(variable *clientside.Bool) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("checked", variable)),
			client.On("input", variable.SetTo(js.Bool{Value: html.Element(c).Get("checked")})),
		)
	})
}

//PartialUpdate renders the checkbox to be partially checked when the given bool is true.
func PartialUpdate(variable *clientside.Bool) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("indeterminate", variable)),
		)
	})
}

//SetReadOnly sets the textbox to be readonly.
func SetReadOnly() seed.Option {
	return attr.Set("readonly", "")
}

//Focus focuses the textbox.
func Focus(c seed.Seed) js.Script {
	return func(q js.Ctx) {
		q(html.Element(c).Run(`focus`))
	}
}
