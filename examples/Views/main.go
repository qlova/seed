package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/s/button"
	"qlova.org/seed/s/text"
	"qlova.org/seed/script"
	"qlova.org/seed/view"
)

type Main struct{}

func (Main) View(c view.Controller) view.Seed {
	return view.New(
		text.New("This is the main view"),
		button.New("Click to go to another view",
			script.OnClick(c.Goto(Another{})),
		),
	)
}

type Another struct{}

func (Another) View(c view.Controller) view.Seed {
	return view.New(
		text.New("This is another view"),
		button.New("Click to go to the main view",
			script.OnClick(c.Goto(Main{})),
		),
	)
}

func main() {
	app.New("Views", view.Set(Main{})).Launch()
}
