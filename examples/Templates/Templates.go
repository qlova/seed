package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/template"

	"github.com/qlova/seeds/text"
)

func main() {
	var App = seed.NewApp()

	var List = template.AddTo(App)
	var Text = text.AddTo(List)
	Text.OnReady(func(q script.Ctx) {
		Text.Ctx(q).SetText(List.Ctx(q).Data().String())
	})

	App.OnReady(func(q script.Ctx) {
		var Array = q.Strings(
			"Apple\n",
			"Apple\n",
			"Apple\n",
			"Apple\n",
		)

		List.Ctx(q).Refresh(Array)
	})

	App.Launch()
}
