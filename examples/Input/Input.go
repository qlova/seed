package main

import (
	"fmt"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/button"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/text"
	"qlova.org/seed/new/textbox"
	"qlova.org/seed/set"
	"qlova.org/seed/set/change"
	"qlova.org/seed/use/css/units/vh"
)

func someServerSideFunction(name string) string {
	message := fmt.Sprintf("Hello %s\n", name)
	fmt.Printf("Created Message: %s\n", message)
	return message
}

func main() {

	Name := &clientside.String{}
	FullName := &clientside.String{}

	app.New("Input Example",
		column.New(
			row.New(
				set.Hidden(),
				set.MarginBottom(vh.New(1)),
				text.New(text.SetStringTo(FullName)),
				change.When(FullName, set.Visible()),
			),
			row.New(
				set.MarginBottom(vh.New(1)),
				text.New(text.Set("What is your name?"),
				),
			),
			row.New(
				textbox.New(
					set.MarginBottom(vh.New(1)),
					textbox.SetPlaceholder("Enter your name"),
					textbox.Update(Name),
				),
			),
			row.New(
				button.New(text.Set("Submit"), client.OnClick(client.Go(func(name string) client.Script {
					newName := someServerSideFunction(name)
					return client.NewScript(
						Name.Set(""),
						FullName.Set(newName),
					)
				}, Name))),
			),
		),
	).Launch()
}
