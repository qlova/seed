package template

import (
	"qlova.org/seed"
	"qlova.org/seed/web/html"
)

//New returns a new HTML template element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("template"), seed.Options(options))
}
