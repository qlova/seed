package main

import "github.com/qlova/seed"
import "github.com/qlova/seeds/text"
import "github.com/qlova/seeds/expander"

func main() {
	var App = seed.NewApp("Hello World")

	expander.AddTo(App)
	text.AddTo(App, "Hello World")
	expander.AddTo(App)

	App.Launch()
}
