package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/s/repeater"
	"qlova.org/seed/s/text"
)

var strings = []string{"a", "b", "c"}

func main() {
	app.New("Repeater",
		repeater.New(strings, repeater.Do(func(c repeater.Seed) {
			c.With(text.New(c.Data.String()))
		})),
	).Launch()
}
