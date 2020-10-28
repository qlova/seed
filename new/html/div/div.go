package div

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

//New returns a new HTML div element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("div"), seed.Options(options))
}
