package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/js/console"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/script"
)

func main() {
	app.New("Script",
		button.New("Click me!",
			script.OnClick(func(ctx script.Ctx) {
				q := js.NewCtx(ctx)

				var s = js.NewString("Hello World!")

				q(console.Log(s))
			}),
		),
	).Launch()
}
