package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/style/align"

	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/text"
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
