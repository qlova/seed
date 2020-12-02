package passwordbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"

	"qlova.org/seed/new/textbox"
)

//New returns a new passwordbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "password"), seed.Options(options))
}

//Update updates the given variable whenever the textbox text is modified.
func Update(variable *clientside.Secret) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("input", variable.SetTo(js.String{Value: html.Element(c).Get("value")})),
		)
	})
}
