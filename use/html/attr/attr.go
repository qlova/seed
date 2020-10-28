package attr

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/html"
)

//Set returns an option setting the HTML element attribute of name to value.
func Set(name string, value string) seed.Option {
	return html.SetAttribute(name, value)
}

//SetTo returns an option setting the HTML element attribute of name to value.
func SetTo(name string, value client.String) seed.Option {
	return html.SetAttributeTo(name, value)
}
