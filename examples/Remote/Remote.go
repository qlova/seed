package main

import "github.com/qlova/seed"
import . "github.com/qlova/script"

func Callback() string {
	return "You clicked me!"
}

func main() {
	var App = seed.Button()
	App.SetText("Click me!")
	
	App.OnClick(func(q seed.Script) {
		q.Get(App).SetText(q.Call(Callback).(String))
	})
	
	App.Launch()
}
