package textarea

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//New returns a new HTML input element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("textarea"), seed.Options(options))
}
