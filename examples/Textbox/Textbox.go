package main

import (
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/state"

	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/s/textbox"
)

func main() {
	Input := state.NewString("", state.Session())

	app.New("Textbox",
		textbox.Var(Input),
		text.Var(state.Sprintf(`Your input: %v`, Input)),
	).Launch()
}
