package a

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

//New returns a new HTML a element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("a"), seed.Options(options))
}
