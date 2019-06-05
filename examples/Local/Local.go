package main

import "github.com/qlova/seed"
import "github.com/qlova/seed/widgets/button"

func main() {
	var App = seed.NewApp("Clientside Code")

	var Button = button.AddTo(App, "Click me!")
	Button.OnClick(func(q seed.Script) {
		Button.Script(q).SetText(q.String("You clicked me!"))
	})

	App.Launch()
}
