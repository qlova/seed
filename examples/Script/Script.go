package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/script"
)

func main() {
	app.New("Script",
		button.New("Click me!",
			script.OnClick(func(q script.Ctx) {
				q.PrintL("Hello World")
			}),
		),
	).Launch()
}
