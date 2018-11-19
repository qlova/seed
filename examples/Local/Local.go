package main

import "github.com/qlova/seed"

func main() {
	var App = seed.Button()
	App.SetText("Click me!")
	
	App.OnClick(func(q seed.Script) {
		q.Get(App).SetText(q.String("You clicked me!"))
	})
	
	App.Launch()
}
