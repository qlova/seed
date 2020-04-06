package main

import (
	"github.com/qlova/seed/app"

	"github.com/qlova/seed/js"
	"github.com/qlova/seed/js/console"
	"github.com/qlova/seed/js/window"

	"github.com/qlova/seed/script"
	"github.com/qlova/seed/script/echo"

	"github.com/qlova/seed/s/button"
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
