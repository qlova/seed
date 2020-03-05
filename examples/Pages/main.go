package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/script"
)

type Home struct{}

func (Home) Page(p page.Seed) {
	p.Add(
		text.New("This is the homepage"),
		button.New("Click to go to another page",
			script.OnClick(p.Goto(Other{})),
		),
	)
}

type Other struct{}

func (Other) Page(p page.Seed) {
	p.Add(
		text.New("This is the other page"),
		button.New("Click to go to the homepage",
			script.OnClick(p.Goto(Home{})),
		),
	)
}

func main() {
	app.New("Pages", page.Set(Home{})).Launch()
}
