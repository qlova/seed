package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/js/console"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/image"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/font"
)

func main() {
	app.New("Script",
		button.New("Click me!",

			style.SetHidden(),
			css.SetDisplay(css.None),
			css.Set("display", "none"),

			script.OnClick(func(q script.Ctx) {
				console.Log(js.NewString(`Hello World`))
				q(`console.log("Hello World")`)
				q.Print(`Hello World`),
			}),
		),

		image.New(`img.png`),
	).Launch()
}
