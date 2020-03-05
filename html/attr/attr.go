package attr

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

//Set returns an option setting the HTML element attribute of name to value.
func Set(name, value string) seed.Option {
	return html.SetAttribute(name, value)
}
