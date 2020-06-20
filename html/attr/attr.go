package attr

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
)

//Set returns an option setting the HTML element attribute of name to value.
func Set(name, value string) seed.Option {
	return html.SetAttribute(name, value)
}
