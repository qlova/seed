package iframe

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/sum"
)

//New returns a new HTML input element.
func New(src sum.String, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("iframe"),
		attr.Set("src", src),
		seed.Options(options),
	)
}
