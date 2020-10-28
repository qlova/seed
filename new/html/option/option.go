package option

import (
	"qlova.org/seed"
	"qlova.org/seed/web/html"
)

//New returns a new HTML option element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("option"),

		seed.Options(options),
	)
}
