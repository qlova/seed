package main

import (
	"qlova.org/seed"
	"qlova.org/seed/app"
	"qlova.org/seed/css"
	"qlova.org/seed/js"
	"qlova.org/seed/js/console"
	"qlova.org/seed/s/button"
	"qlova.org/seed/s/image"
	"qlova.org/seed/script"
	"qlova.org/seed/style"
	"qlova.org/seed/style/font"
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
