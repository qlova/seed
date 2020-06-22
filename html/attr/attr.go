package attr

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
	"qlova.org/seed/sum"
)

//Set returns an option setting the HTML element attribute of name to value.
func Set(name string, value sum.String) seed.Option {
	return html.SetAttribute(name, value)
}
