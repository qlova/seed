package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/tween"

	"github.com/qlova/seed/s/row"
	"github.com/qlova/seed/s/text"
)

func main() {
	DisplayAsColumn := state.New()

	app.New("Tween",
		row.New(
			text.New("A", tween.This()),
			text.New("B", tween.This()),
			text.New("C", tween.This()),

			DisplayAsColumn.If(
				style.SetColumn(),
			),
		),

		script.OnClick(tween.Tween(DisplayAsColumn.Toggle())),
	).Launch()
}
