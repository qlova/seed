package main

import (
	"qlova.org/seed"
	"qlova.org/seed/app"
	"qlova.org/seed/s/button"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/style"
	"qlova.org/seed/user"
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
