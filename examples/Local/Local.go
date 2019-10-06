package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/button"
)

func main() {
	var App = seed.NewApp("Clientside Code")

	var Button = button.AddTo(App, "Click me!")
	Button.OnClick(func(q script.Ctx) {
		Button.Ctx(q).SetText(q.String("You clicked me!"))
	})

	App.Launch()
}
