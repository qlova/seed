package main

import (
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/script/global"

	"github.com/qlova/seed"
	"github.com/qlova/seeds/button"
	"github.com/qlova/seeds/header"
	"github.com/qlova/seeds/text"
)

//Counter is the variable storing the current count.
var Counter = global.NewInt()

func main() {
	var App = seed.NewApp()

	header.AddTo(App, "Counter")
	text.AddTo(App).SetTextf(`Current count: "%v"`, Counter)

	button.AddTo(App, "Click me!").OnClick(func(q script.Ctx) {
		Counter.PlusPlus(q)
	})
	button.AddTo(App, "Reset").OnClick(func(q script.Ctx) {
		Counter.Set(q, q.Int(0))
	})

	App.Launch()
}
