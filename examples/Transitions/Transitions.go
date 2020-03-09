package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/page/transition"
	"github.com/qlova/seed/script"

	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/text"
)

type Home struct{}

func (Home) Page(p page.Seed) {
	p.Add(
		transition.Fade(),

		text.New("This is the homepage"),
		button.New("Click to go to another page",
			script.OnClick(p.Goto(Other{})),
		),
	)
}

type Other struct{}

func (Other) Page(p page.Seed) {
	p.Add(
		transition.Fade(),

		text.New("This is the other page"),
		button.New("Click to go to the homepage",
			script.OnClick(p.Goto(Home{})),
		),
	)
}

func main() {
	app.New("Transitions", page.Set(Home{})).Launch()
}
