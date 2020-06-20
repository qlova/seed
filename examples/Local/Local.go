package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/client"
	"qlova.org/seed/s/button"
	//"qlova.org/seed/s/button"
)

func main() {
	Message := client.NewStringVar("Click me!")

	app.New("Local Code",
		button.New(Message,
			client.OnClick(Message.Set("You Clicked me!")),
		),
	).Launch()
}
