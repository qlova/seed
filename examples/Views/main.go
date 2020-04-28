package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/view"
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
