package emailbox

import (
	"qlova.org/seed"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/sum"

	"qlova.org/seed/s/textbox"
)

//New returns a new emailbox widget.
func New(sync sum.String, options ...seed.Option) seed.Seed {
	return textbox.New(sync, attr.Set("type", "email"), seed.Options(options))
}
