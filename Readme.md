# ![logo](media/logo.png) Qlovaseed 

An eco-friendly, cross-platform application framework for Go.

Hello World:
```
	package main

	import "github.com/qlova/seed"

	func main() {
		var App = seed.New()
		App.SetName("Hello World")
		App.SetText("Hello World")
		App.Launch()
	}
```
