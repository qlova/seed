package table

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

//New returns an HTML 'table' element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("table"), seed.Options(options))
}
