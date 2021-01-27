package main

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/button"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/text"
	"qlova.org/seed/new/view"
	"qlova.org/seed/set"
	"qlova.org/seed/use/css/units/vh"
)

type ViewOne struct {
}

func (f ViewOne) View(c view.Controller) seed.Seed {
	return view.New(
		column.New(
			row.New(
				text.New(text.Set("View Number One")),
			),
			row.New(
				button.New(text.Set("Next"), client.OnClick(c.Next())),
			),
		),
	)
}

type ViewTwo struct {
}

func (f ViewTwo) View(c view.Controller) seed.Seed {
	return view.New(
		column.New(
			row.New(
				text.New(text.Set("View Number Two")),
			),
			row.New(
				button.New(text.Set("Back"), client.OnClick(c.Back())),
			),
		),
	)
}

type StandaloneView struct {
}

func (f StandaloneView) View(_ view.Controller) seed.Seed {
	return view.New(
		column.New(
			row.New(
				text.New(text.Set("View not in a list")),
			),
		),
	)
}

func main() {
	app.New("Views Example",
		column.New(
			row.New(
				set.MarginBottom(vh.New(2)),
				text.New(
					text.Set("Welcome to our Views Example"),
				),
			),
			row.New(
				set.MarginBottom(vh.New(2)),
				view.List(ViewOne{}, ViewTwo{}),
			),
			row.New(
				set.MarginBottom(vh.New(2)),
				text.New(
					text.Set("Aren't they great?"),
				),
			),
			row.New(
				view.Set(StandaloneView{}),
			),
		),
	).Launch()
}
