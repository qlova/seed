package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/button"
)

//Import a widget to use it, a list of widgets can be found in the widgets directory.

func main() {
	var App = seed.NewApp("My App")

	//In order to add a widget to your app, or container, use the package's AddTo method.
	var ClientPowered = button.AddTo(App, "My callback runs on the client")

	ClientPowered.OnClick(func(q script.Ctx) {
		ClientPowered.Ctx(q).SetText(q.String("You clicked me!"))
	})

	var ServerPowered = button.AddTo(App, "My callback runs on the server")

	//You can style widgets with methods of the style package.
	ServerPowered.SetColor(seed.RGB(100, 100, 0))

	ServerPowered.OnClick(seed.Go(func(user seed.User) {
		ServerPowered.For(user).SetText("You clicked me!")
	}))

	App.Launch()
}
