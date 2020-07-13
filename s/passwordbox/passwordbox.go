package passwordbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/script"

	"qlova.org/seed/s/textbox"
)

//New returns a new passwordbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New( attr.Set("type", "password"), seed.Options(options))
}

//Update updates the given variable whenever the textbox text is modified.
func Update(variable *clientside.Secret) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			script.On("input", variable.SetTo(js.String{Value: script.Element(c).Get("value")})),
		)
	})
}
