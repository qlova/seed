package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/image"
)

func main() {
	app.New("Image",
		image.New("logo.svg"),
	).Launch()
}
