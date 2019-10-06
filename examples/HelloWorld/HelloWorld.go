package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/expander"
	"github.com/qlova/seeds/text"
)

func main() {
	var App = seed.NewApp("Hello World")

	expander.AddTo(App)
	text.AddTo(App, "Hello World")
	expander.AddTo(App)

	App.Launch()
}
