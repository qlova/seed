package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/script/global"

	"github.com/qlova/seeds/button"
	"github.com/qlova/seeds/text"
)

var Time = global.NewString()

func main() {
	var App = seed.NewApp("Dynamic")

	for i := 0; i < 3; i++ {
		text.AddTo(App).SetDynamicText(Time)
	}

	button.AddTo(App, "Click Me!").OnClick(func(q script.Ctx) {
		Time.Set(q, q.Time.Now().String())
	})

	App.Launch()
}
