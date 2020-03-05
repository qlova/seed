package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/repeater"
	"github.com/qlova/seed/s/text"
)

var strings = []string{"a", "b", "c"}

func main() {
	app.New("Repeater",
		repeater.New(strings, seed.Do(func(c seed.Seed) {
			c.Add(text.New(repeater.Data(c).String()))
		})),
	).Launch()
}
