package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/style/font"
)

var LoggedIn = state.New()

func main() {
	app.New("HelloWorld",

		LoggedIn.If(
			text.New("You are logged in",

				font.SetColor(seed.Green),
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
