package input

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//New returns a new HTML input element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("input").And(options...))
}
