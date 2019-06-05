package main

import "github.com/qlova/seed"
import "github.com/qlova/seed/widgets/button"

func main() {
	var App = seed.NewApp("Remote Code")

	var Button = button.AddTo(App, "Click me!")
	Button.OnClick(seed.Go(func(user seed.User) {
		Button.For(user).SetText("You clicked me!")
	}))

	App.Launch()
}
