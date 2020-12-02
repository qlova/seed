package datebox

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html/attr"

	"qlova.org/seed/new/textbox"
)

//New returns a new datebox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "date"), seed.Options(options))
}
