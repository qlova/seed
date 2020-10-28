package input

import (
	"qlova.org/seed"
	"qlova.org/seed/web/html"
)

//New returns a new HTML input element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("input"), seed.Options(options))
}
