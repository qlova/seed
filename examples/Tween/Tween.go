package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/style"
	"qlova.org/seed/tween"

	"qlova.org/seed/s/row"
	"qlova.org/seed/s/text"
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
