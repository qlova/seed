package textbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/script"

	"qlova.org/seed/s/html/input"
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
			client.On("render", script.Element(c).Set("value", variable)),
			client.On("input", variable.SetTo(js.String{Value: script.Element(c).Get("value")})),
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
func Focus(c seed.Seed) script.Script {
	return func(q script.Ctx) {
		q(script.Element(c).Run(`focus`))
	}
}
