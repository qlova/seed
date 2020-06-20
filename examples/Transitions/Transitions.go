package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/js"
	"qlova.org/seed/page"
	"qlova.org/seed/page/transition"
	"qlova.org/seed/script"

	"qlova.org/seed/s/button"
	"qlova.org/seed/s/text"
)

type Home struct {
	Content js.String `url:"content"`
}

func (h Home) Page(r page.Router) page.Seed {
	return page.New(
		page.SetTitle(""),
		page.SetPath("/home"),

		transition.Fade(),

		text.Var(h.Content),
		button.New("Click to go to another page",
			script.OnClick(r.Goto(Other{})),
		),
	)
}

type Other struct{}

func (Other) Page(r page.Router) page.Seed {
	return page.New(
		transition.Fade(),

		text.New("This is the other page"),
		button.New("Click to go to the homepage",
			script.OnClick(r.Goto(Home{
				Content: js.NewString(`This is the homepage`),
			})),
		),
	)
}

func main() {
	app.New("Transitions", app.SetPage(Home{})).Launch()
}
