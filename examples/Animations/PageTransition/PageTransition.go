package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/button"
)

func main() {
	var App = seed.NewApp()

	var Home = App.NewPage()
	Home.SetColor(seed.Green)
	Home.SetTransition(seed.FlipOut)

	var Page = App.NewPage()
	Page.SetColor(seed.Red)
	Page.SetTransition(seed.FadeIn)

	button.AddTo(Home, "Click Me!").OnClickGoto(Page)
	button.AddTo(Page, "Click Me!").OnClickGoto(Home)

	App.SetPage(Home)
	App.Launch()
}
