package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/s/image"
)

func main() {
	app.New("Image",
		image.New("logo.svg"),
	).Launch()
}
