package monthbox

import (
	"qlova.org/seed"
	"qlova.org/seed/html/attr"

	"qlova.org/seed/s/textbox"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(nil, attr.Set("type", "month"), seed.Options(options))
}
