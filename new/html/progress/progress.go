package progress

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

//New returns a new HTML progress element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("progress"),

		seed.Options(options),
	)
}
