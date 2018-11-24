# ![logo](media/logo.png) Qlovaseed 

Qlovaseed allows cross-platform progressive web applications to be written in pure Go, without touching Html, Css or Javascript.
Currently very unstable and in Heavy Development.

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
