package main

import "time"
import "github.com/qlova/seed"

func Time(client seed.Client) {
	client.WriteString(time.Now().Format("3:04:05PM"))
}

func main() {
	var App = seed.New()
	App.SetName("Dynamic Time")

	App.SetDynamicText(Time)

	App.Launch()
}
