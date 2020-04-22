package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/js/console"
	"github.com/qlova/seed/s/passwordbox"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state/secret"
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
