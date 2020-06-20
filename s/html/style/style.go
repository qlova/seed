package style

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
)

//New returns an HTML style element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("style"), seed.Options(options))
}
