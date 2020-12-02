package canvas

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

//New returns a new HTML canvas element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("canvas"), seed.Options(options))
}
