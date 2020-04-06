package canvas

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//New returns a new HTML canvas element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("canvas").And(options...))
}
