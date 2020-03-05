package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/text"
)

func main() {
	app.New("HelloWorld", text.New("Hello World")).Launch()
}
