package main

import "os"
import "io/ioutil"
import "github.com/qlova/seed" 
import "github.com/qlova/seed/widgets/editor"

func main() {
	var Editor = editor.New()
	Editor.SetName("Text Editor")
	
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			Editor.OnReady(func(q seed.Script) {
				q.Alert(q.String(err.Error()))
			})
		} else {
			Editor.SetText(string(data))
		}
	}
	
	Editor.Launch()
}
