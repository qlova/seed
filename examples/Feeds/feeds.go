package main

import (
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/button"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/text"
	"qlova.org/seed/new/textbox"
	"qlova.org/seed/set"
	"qlova.org/seed/use/css/units/vh"
	"qlova.org/seed/use/js"
)

func main() {
	newFruit := &clientside.String{}

	values := []string{"Apples", "Pears"}
	names := feed.With(feed.Go(func() []string {
		return values
	}))

	app.New("Feeds",
		column.New(
			row.New(
				set.MarginBottom(vh.New(1)),
				text.New(text.Set("Fruit List"), text.SetSize(vh.New(3))),
			),
			row.New(
				names.New(
					row.New(
						text.New(
							text.SetStringTo(js.String{Value: names.Data.Value}),
						),
					),
				),
				client.OnLoad(names.Refresh()),
				set.MarginBottom(vh.New(1)),
			),
			row.New(
				textbox.New(textbox.SetPlaceholder("Enter a fruit"), textbox.Update(newFruit)),
				button.New(text.Set("Add Fruit"), client.OnClick(client.Go(func(_newFruit string) client.Script {
					if _newFruit == "" {
						return client.NewScript()
					}
					values = append(values, _newFruit)
					return client.NewScript(newFruit.Set(""), names.Refresh())
				}, newFruit))),
			),
		),
	).Launch()
}
