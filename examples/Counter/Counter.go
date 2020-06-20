package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/style/align"

	"qlova.org/seed/s/button"
	"qlova.org/seed/s/text"
)

func main() {
	count := state.NewInt(0)

	app.New("Counter",
		text.New("Counter", align.Center()),
		text.Var(state.Sprintf(`Current count: %v`, count), align.Center()),

		button.New("Click me!", script.OnClick(count.Increment())),
		button.New("Reset", script.OnClick(count.SetL(0))),
	).Launch()
}
