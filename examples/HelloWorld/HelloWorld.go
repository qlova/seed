package main

import "github.com/qlova/seed"

func main() {
	var App = seed.NewApp()
	App.SetName("Hello World")
	App.SetText("Hello World")
	App.Launch()
}
