package app

import "bytes"
import "fmt"

type Script struct {
	element *App
	script bytes.Buffer
}

func (script *Script) Get(app *App) *Script {
	script.element = app
	return script
}

func (script *Script) Set(property, value string) {
	script.script.WriteString(`document.getElementById("`)
	script.script.WriteString(fmt.Sprint(script.element.id))
	script.script.WriteString(`").`)
	script.script.WriteString(property)
	script.script.WriteByte('=')
	script.script.WriteString(value)
	script.script.WriteByte(';')
}

func (script *Script) LogString(msg string) {
	script.script.WriteString(`console.log("`)
	script.script.WriteString(msg)
	script.script.WriteString(`"`)
}

func (script *Script) SetHidden() {
	script.Set("style.display", `"none"`)	
}

func (script *Script) SetVisible() {
	script.Set("style.display", `"initial"`)
}

//Add text, html or whatever!
func (script *Script) SetPage(page *App) {
	for _, child := range script.element.children {
		if child.page {
			if child == page {
				script.Get(child).SetVisible()
			} else {
				script.Get(child).SetHidden()
			}
		}
	}
}
