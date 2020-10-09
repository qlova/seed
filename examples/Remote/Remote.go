package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/s/button"
	"qlova.org/seed/s/text"
	"qlova.tech/rgb"
)

func main() {

	Message := &clientside.String{Value: "Click me!"}

	app.New("Remote Code",
		button.New(text.SetStringTo(Message),

			text.SetColor(rgb.Red),

			client.OnClick(Test(clientargs.Go, Message)),
		),
	).Launch()
}
