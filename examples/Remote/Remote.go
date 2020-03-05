package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/user"
)

func main() {

	Message := state.NewString("Click me!")

	app.New("Remote Code",
		button.Var(Message,

			style.SetTextColor(seed.Red),

			script.OnClick(script.Go(func(u user.Ctx) {
				Message.For(u).Set("You clicked me!")
			})),
		),
	).Launch()
}
