package iframe

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
	"qlova.org/seed/state"
)

//New returns a new HTML input element.
func New(src state.AnyString, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("iframe"),
		state.SetSource(src),
		seed.Options(options),
	)
}
