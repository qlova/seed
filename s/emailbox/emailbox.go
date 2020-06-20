package emailbox

import (
	"qlova.org/seed"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/state"

	"qlova.org/seed/s/textbox"
)

//New returns a new emailbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "email"), seed.Options(options))
}

//Var returns an emailbox with a variable state argument that will be synced with the value of this emailbox.
func Var(text state.String, options ...seed.Option) seed.Seed {
	return textbox.Var(text, attr.Set("type", "email"), seed.Options(options))
}
