package main

import (
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/button"
	"qlova.org/seed/new/text"
	"qlova.tech/rgb"
)

func main() {

	Message := &clientside.String{Value: "Click me!"}

	app.New("Remote Code",
		button.New(text.SetStringTo(Message),

			text.SetColor(rgb.Red),

			client.OnClick(Message.Set("hello")),
		),
	).Launch()
}
