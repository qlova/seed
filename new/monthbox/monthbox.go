package monthbox

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html/attr"

	"qlova.org/seed/new/textbox"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "month"), seed.Options(options))
}
