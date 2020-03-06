package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/style"
)

func main() {
	Message := state.NewString("Click me!")

	app.New("Local Code",
		button.Var(Message,

			style.SetTextColor(seed.Red),

			script.OnClick(func(q script.Ctx) {
				Message.SetL("You clicked me!")(q)
			}),
		),
	).Launch()
}
