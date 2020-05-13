package iframe

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/state"
)

//New returns a new HTML input element.
func New(src state.AnyString, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("iframe"),
		state.SetSource(src),
		seed.Options(options),
	)
}
