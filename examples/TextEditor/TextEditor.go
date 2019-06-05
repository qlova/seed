package main

import "os"
import "io/ioutil"
import "github.com/qlova/seed"
import "github.com/qlova/seed/widgets/editor"

//A basic text editor without the ability to save files.
func main() {
	var App = seed.NewApp("Text Editor")
	App.SetSize(100, 100)

	var Editor = editor.AddTo(App)
	Editor.SetSize(100, 100)

	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			Editor.OnReady(func(q seed.Script) {
				q.Alert(q.String(err.Error()))
			})
		} else {
			Editor.SetContent(string(data))
		}
	}

	App.Launch()
}
