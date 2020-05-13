package form

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//New returns a new HTML form element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("form"),
		seed.Options(options),
	)
}
