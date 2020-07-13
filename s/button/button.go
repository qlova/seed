package button

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
)

//New returns a button with the given label.
//label can be nil, a string, Stringable or client.String
//anything else is passed to fmt.Sprint and then treated as a string.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("button"),
		attr.Set("type", "button"),

		seed.Options(options),
	)
}
