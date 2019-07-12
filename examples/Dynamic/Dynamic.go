package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	
	"github.com/qlova/seeds/text"
	"github.com/qlova/seeds/button"
)

var Time = script.NewString()

func main() {
	var App = seed.NewApp("Dynamic")

	for i := 0; i < 3; i++ {
		text.AddTo(App).SetDynamicText(Time)
	}

	button.AddTo(App, "Click Me!").OnClick(func(q seed.Script) {
		q.Set(Time, q.Time.Now().String())
	})

	App.Launch()
}
