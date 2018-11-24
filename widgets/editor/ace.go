package editor

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

func init() {
	seed.Embed("/ace.js", []byte(Javascript))
}

//Returns a full-featured text editor with line numbers and optional syntax highlighting.
func New() seed.Seed {
	editor := seed.New()
	editor.SetTag("pre")
	
	editor.Require("ace.js")
	
	editor.OnReady(func(q seed.Script) {
		q.Javascript(`let editor = ace.edit("`+editor.ID()+`"); editor.setShowPrintMargin(false); document.getElementById("`+editor.ID()+`").editor = editor;`)
	})
	
	return editor
}

type Editor script.Seed

func (editor Editor) Open(f script.File) {
	script.Seed(editor).Javascript(`var reader = new FileReader(); reader.onload = function(e) { var data = e.target.result; get("`+editor.ID+`").editor.setValue(data); }; reader.readAsText(`+string(f)+`);`)
}
