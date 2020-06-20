package main

import (
	"qlova.org/seed"
	"qlova.org/seed/app"
	"qlova.org/seed/s/button"
	"qlova.org/seed/s/text"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/style/font"
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
