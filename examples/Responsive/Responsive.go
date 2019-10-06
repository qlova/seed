package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/expander"
	"github.com/qlova/seeds/text"
)

func main() {
	var App = seed.NewApp("Responsive")

	expander.AddTo(App)
	var Tiny = text.AddTo(App, "Tiny")
	{
		Tiny.SetHidden()
		Tiny.Tiny.SetVisible()
	}
	var Small = text.AddTo(App, "Small")
	{
		Small.SetHidden()
		Small.Small.SetVisible()
	}
	var Medium = text.AddTo(App, "Medium")
	{
		Medium.SetHidden()
		Medium.Medium.SetVisible()
	}
	var Large = text.AddTo(App, "Large")
	{
		Large.SetHidden()
		Large.Large.SetVisible()
	}
	var Huge = text.AddTo(App, "Huge")
	{
		Huge.SetHidden()
		Huge.Huge.SetVisible()
	}
	var Portrait = text.AddTo(App, "Portrait")
	{
		Portrait.SetHidden()
		Portrait.Portrait.SetVisible()
	}
	var Landscape = text.AddTo(App, "Landscape")
	{
		Landscape.SetHidden()
		Landscape.Landscape.SetVisible()
	}
	expander.AddTo(App)

	App.Launch()
}
