package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/s/text"
)

func main() {
	app.New("Hello World",
		text.Set("Hello World"),
	).Launch()
}
