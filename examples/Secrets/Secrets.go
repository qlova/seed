package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/js/console"
	"qlova.org/seed/s/passwordbox"
	"qlova.org/seed/script"
	"qlova.org/seed/state/secret"
)

func main() {
	var Password = secret.New("secrets")

	app.New("Secrets",
		passwordbox.Var(Password,
			script.OnEnter(func(q script.Ctx) {
				q(console.Log(Password))
			}),
		),
	).Launch()
}
