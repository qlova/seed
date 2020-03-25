package monthbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/state"

	"github.com/qlova/seed/s/textbox"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "month").And(options...))
}

//Var returns text with a variable text argument.
func Var(text state.String, options ...seed.Option) seed.Seed {
	return textbox.Var(text, attr.Set("type", "month").And(options...))
}
