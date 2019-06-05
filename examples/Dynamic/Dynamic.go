package main

import "time"
import "github.com/qlova/seed"

func Time(client seed.User) {
	client.WriteString(time.Now().Format("3:04:05PM"))
}

func main() {
	var App = seed.NewApp("Dynamic Time")

	var Text = seed.AddTo(App)
	Text.SetDynamicText(Time)

	App.Launch()
}
