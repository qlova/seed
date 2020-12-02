package iframe

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
)

//New returns a new HTML input element.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("iframe"),

		seed.Options(options),
	)
}

func Set(src string) seed.Option {
	return attr.Set("src", src)
}

func SetTo(src client.String) seed.Option {
	return attr.SetTo("src", src)
}
