package datalist

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html"
)

//New returns a new HTML datalist element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("datalist"),

		seed.Options(options),
	)
}
