package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/page/transition"
	"github.com/qlova/seed/script"

	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/text"
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
