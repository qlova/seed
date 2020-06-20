package monthbox

import (
	"qlova.org/seed"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/state"

	"qlova.org/seed/s/textbox"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "month"), seed.Options(options))
}

//Var returns text with a variable text argument.
func Var(text state.String, options ...seed.Option) seed.Seed {
	return textbox.Var(text, attr.Set("type", "month"), seed.Options(options))
}
