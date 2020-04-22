package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/script"
)

type Home struct{}

func (Home) Page(r page.Router) page.Seed {
	return page.New(
		text.New("This is the homepage"),
		button.New("Click to go to another page",
			script.OnClick(r.Goto(Other{})),
		),
	)
}

type Other struct{}

func (Other) Page(r page.Router) page.Seed {
	return page.New(
		text.New("This is the other page"),
		button.New("Click to go to the homepage",
			script.OnClick(r.Goto(Home{})),
		),
	)
}

func main() {
	app.New("Pages", app.SetPage(Home{})).Launch()
}
