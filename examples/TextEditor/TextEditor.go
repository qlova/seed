package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seeds/editor"
)

//A basic text editor without the ability to save files.
func main() {
	var App = seed.NewApp("Text Editor")
	App.SetSize(100, 100)

	var Editor = editor.AddTo(App)
	Editor.SetSize(100, 100)

	App.Launch()
}
