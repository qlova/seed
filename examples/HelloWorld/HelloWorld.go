package main

import (
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/text"
)

func main() {
	app.New("Hello World",
		text.Set("Hello World"),
	).Launch()
}
