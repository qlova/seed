// +build !generate,!wasm

package main

import (
	"qlova.org/seed/client"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/use/js"

	"dating"
	"dating/ui"
)

func main() {

	app.New("DatingApp",
		row.Set(),

		client.OnLoad(
			client.Run(dating.LoadCustom, js.Func("window.localStorage.getItem").Call(client.NewString("custom.dates"))),
		),

		ui.NewSidebar(),
		page.AddPages(ui.PopularPage{}, ui.CustomPage{}, ui.AddPage{}),
		page.Set(ui.PopularPage{}),
	).Launch()
}
