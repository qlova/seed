package emailbox

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html/attr"

	"qlova.org/seed/new/textbox"
)

//New returns a new emailbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "email"), seed.Options(options))
}
