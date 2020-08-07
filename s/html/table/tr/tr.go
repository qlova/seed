package tr

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
)

//New returns an HTML 'tr' element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("tr"), seed.Options(options))
}
