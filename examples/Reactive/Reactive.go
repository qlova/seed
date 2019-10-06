package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seeds/button"
	"github.com/qlova/seeds/text"
)

var ButtonState = seed.NewState()

func main() {
	var App = seed.NewApp("Reactive")

	for i := 0; i < 3; i++ {
		var Text = text.AddTo(App)
		Text.When(ButtonState, func(q script.Ctx) {
			Text.Ctx(q).SetText(q.String("We are in button state"))
		})
		Text.When(ButtonState.Not(), func(q script.Ctx) {
			Text.Ctx(q).SetText(q.String("We are not in button state"))
		})
	}

	button.AddTo(App, "Click Me!").OnClickToggleState(ButtonState)

	App.Launch()
}
