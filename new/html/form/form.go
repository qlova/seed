package form

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

//New returns a new HTML form element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("form"),
		seed.Options(options),
	)
}
