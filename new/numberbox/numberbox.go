package numberbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/new/html/input"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(
		attr.Set("type", "number"),
		attr.Set("inputmode", "numeric"),
		attr.Set("pattern", "[0-9]"),
		seed.Options(options),
	)
}

//Update syncs the given variable with the numberbox's value.
func Update(variable *clientside.Float64) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("value", variable)),
			client.On("input", variable.SetTo(js.Number{js.NewValue("+%v", html.Element(c).Get("value"))})),
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
