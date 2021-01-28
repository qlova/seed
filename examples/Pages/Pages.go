package main

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/button"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/text"
	"qlova.org/seed/use/css"
)

type HomePage struct {
}

func (p HomePage) Page(r page.Router) seed.Seed {
	return page.New(
		column.New(
			css.Set("align-items", "center"),
			row.New(
				text.New(text.SetString("Welcome to our Home Page!")),
			),
			row.New(
				button.New(
					text.Set("Learn About Us"),
					client.OnClick(r.Goto(AboutPage{}),
					),
				),
			),
		),
	)
}

type AboutPage struct {
}

func (p AboutPage) Page(r page.Router) seed.Seed {
	return page.New(
		column.New(
			css.Set("align-items", "center"),
			row.New(
				text.New(text.SetString("All About Us!")),
			),
			row.New(
				button.New(
					text.Set("Home"),
					client.OnClick(r.Goto(HomePage{}),
					),
				),
			),
		),
	)
}

func main() {
	app.New("Pages Example",
		page.AddPages(HomePage{}, AboutPage{}),
		page.Set(HomePage{}),
	).Launch()
}
