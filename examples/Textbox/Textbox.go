package main

import (
	"qlova.org/seed/app"
	"qlova.org/seed/state"

	"qlova.org/seed/s/text"
	"qlova.org/seed/s/textbox"
)

func main() {
	Input := state.NewString("", state.Session())

	app.New("Textbox",
		textbox.Var(Input),
		text.Var(state.Sprintf(`Your input: %v`, Input)),
	).Launch()
}
