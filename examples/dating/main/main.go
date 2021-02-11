// +build !generate,!wasm

package main

import (
	"fmt"

	"qlova.org/seed/client"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/set"
	"qlova.org/seed/use/js"
	"qlova.tech/rgb"

	"dating"
	"dating/ui"
)

func main() {

	var DatingApp = app.New("DatingApp",
		app.SetColor(rgb.Black),

		app.OnUpdateFound(app.Update()),

		row.Set(),
		set.If.Small().Portrait(
			column.Set(),
		),

		client.OnLoad(
			client.Run(dating.LoadCustom, js.Func("window.localStorage.getItem").Call(client.NewString("custom.dates"))),
		),

		ui.NewSidebar(),
		page.AddPages(ui.PopularPage{}, ui.CustomPage{}, ui.AddPage{}),
		page.Set(ui.PopularPage{}),
	)

	if err := DatingApp.Export(); err != nil {
		fmt.Println(err)
	}

	DatingApp.Launch()
}
