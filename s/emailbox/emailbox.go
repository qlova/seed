package emailbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/state"

	"github.com/qlova/seed/s/textbox"
)

//New returns a new emailbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "email").And(options...))
}

//Var returns an emailbox with a variable state argument that will be synced with the value of this emailbox.
func Var(text state.String, options ...seed.Option) seed.Seed {
	return textbox.Var(text, attr.Set("type", "email").And(options...))
}
