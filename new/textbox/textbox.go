package textbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/web/html"
	"qlova.org/seed/web/html/attr"
	"qlova.org/seed/web/js"

	"qlova.org/seed/new/html/input"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(
		seed.Options(options),
	)
}

//Update updates the given variable whenever the textbox text is modified.
func Update(variable *clientside.String) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("value", variable)),
			client.On("input", variable.SetTo(js.String{Value: html.Element(c).Get("value")})),
		)
	})
}

//SetPlaceholder sets the placeholder of the textbox.
func SetPlaceholder(placeholder string) seed.Option {
	return attr.Set("placeholder", placeholder)
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
