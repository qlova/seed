package template

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//New returns a new HTML template element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("template"), seed.Options(options))
}
