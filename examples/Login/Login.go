package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
)

var LoggedIn = state.New()

func main() {
	app.New("HelloWorld",

		LoggedIn.If(
			text.New("You are logged in",

				text.SetColor(seed.Green),
			),
		),

		button.New("Login",

			script.OnClick(LoggedIn.Toggle()),

			LoggedIn.If(
				text.Set("Logout"),
			),
		),
	).Launch()
}
