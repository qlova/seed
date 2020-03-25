package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/repeater"
	"github.com/qlova/seed/s/text"
)

var strings = []string{"a", "b", "c"}

func main() {
	app.New("Repeater",
		repeater.New(strings, repeater.Do(func(c repeater.Seed) {
			c.Add(text.New(c.Data.String()))
		})),
	).Launch()
}
