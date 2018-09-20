package script

import "bytes"
import "github.com/qlova/app/interfaces"
import "github.com/qlova/app/style"
import "strings"

type App struct {
	interfaces.App
	
	script *Script
	style *style.Style
}

func (app *App) GetParent() interfaces.App {
	return app.script.Get(app.App.GetParent())
}

func (app *App) GetChildren() []interfaces.App {
	var children = app.App.GetChildren()
	var result = make([]interfaces.App, len(children))
	for i := range children {
		result[i] = app.script.Get(children[i])
	}
	return result
}

func (app *App) Click() {
	app.script.data.WriteString(`document.getElementById("`)
	app.script.data.WriteString(app.ID())
	app.script.data.WriteString(`").click();`)
}


func (app *App) GetStyle() *style.Style {
	return app.style
}


type Script struct {
	data bytes.Buffer
}

func (script *Script) Bytes() []byte {
	return script.data.Bytes()
}

func (script *Script) Get(app interfaces.App) *App {
	if app == nil {
		return nil
	}
	
	if sa, ok := app.(*App); ok {
		return sa
	}
	
	sa := new(App)
	sa.script = script
	sa.App = app
	sa.style = &style.Style{Css: &scriptCss{script:script, app:sa}}

	return sa
}

func (script *Script) LogString(msg string) {
	script.data.WriteString(`console.log("`)
	script.data.WriteString(msg)
	script.data.WriteString(`"`)
}

type scriptCss struct {
	script *Script
	app *App
}

func (css *scriptCss) Set(property, value string) {
	
	if strings.Contains(property, "-") {
		splits := strings.Split(property, "-")
		
		property = splits[0] + strings.Title(splits[1])
	}
	
	css.script.data.WriteString(`document.getElementById("`)
	css.script.data.WriteString(css.app.ID())
	css.script.data.WriteString(`").style.`)
	css.script.data.WriteString(property)
	css.script.data.WriteByte('=')
	css.script.data.WriteByte('"')
	css.script.data.WriteString(value)
	css.script.data.WriteByte('"')
	css.script.data.WriteByte(';')
}
