package div

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//New returns a new HTML div element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("div").And(options...))
}
