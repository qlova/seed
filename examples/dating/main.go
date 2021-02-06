package main

import (
	"qlova.org/seed/client"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/use/js"
)

func main() {

	app.New("Timers",
		row.Set(),

		client.OnLoad(
			client.Run(LoadCustom, js.Func("window.localStorage.getItem").Call(client.NewString("custom.dates"))),
		),

		NewSidebar(),
		page.AddPages(PopularPage{}, CustomPage{}, AddPage{}),
		page.Set(PopularPage{}),
	).Launch()
}
