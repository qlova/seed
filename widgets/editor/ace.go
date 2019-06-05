package editor

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

func init() {
	seed.Embed("/ace.js", []byte(Javascript))
}

type Widget struct {
	seed.Seed
}

//Returns a full-featured text editor with line numbers and optional syntax highlighting.
func New(text ...string) Widget {
	widget := seed.New()
	widget.SetTag("pre")

	if len(text) > 0 {
		widget.SetText(text[0])
	}

	widget.Require("ace.js")

	widget.OnReady(func(q seed.Script) {
		q.Javascript(`let editor = ace.edit("` + widget.ID() + `"); editor.setShowPrintMargin(false); document.getElementById("` + widget.ID() + `").editor = editor;`)
	})

	return Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, path ...string) Widget {
	var widget = New(path...)
	parent.Root().Add(widget)
	return widget
}

type Editor struct {
	script.Seed
}

func (editor Editor) Open(f script.File) {
	editor.Javascript(`var reader = new FileReader(); reader.onload = function(e) { var data = e.target.result; get("` + editor.ID + `").editor.setValue(data); }; reader.readAsText(` + f.Raw() + `);`)
}
