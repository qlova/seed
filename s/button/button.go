package button

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/sum"
)

//New returns a button with the given label.
//label can be nil, a string, Stringable or client.String
//anything else is passed to fmt.Sprint and then treated as a string.
func New(label sum.String, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("button"),
		attr.Set("type", "button"),

		html.SetInnerText(label),

		seed.Options(options),
	)
}
