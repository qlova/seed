package main

import (
	"qlova.org/seed/app"

	"qlova.org/seed/js"
	"qlova.org/seed/js/console"
	"qlova.org/seed/js/window"

	"qlova.org/seed/script"
	"qlova.org/seed/script/echo"

	"qlova.org/seed/s/button"
)

type Test struct {
	Name string
}

func main() {
	app.New("Objects",
		button.New("Click me!",
			script.OnClick(func(q script.Ctx) {
				var t = js.ValueOf(Test{}).Var(q)

				q(echo.ChangesTo(t, func(e echo.Ctx, t *Test) {
					t.Name = e.String(window.Prompt(q.String("What is your name?")))
				}))

				q(console.Log(t))
			}),
		),
	).Launch()
}
