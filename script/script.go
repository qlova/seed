package script

import "bytes"
import "github.com/qlova/app/interfaces"
import "github.com/qlova/app/style"
import "strings"

type CssWrapper struct {
	app interfaces.App
}

func (c CssWrapper) Set(a, b string) {
	c.app.GetStyle().Css.Set(a, b)
}

type App struct {
	interfaces.App
	style.Style
	
	script *Script
	style *style.Style
	
	query string
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
	app.Run("click");
}

func (app *App) Run(method string) {
	if app.query != "" {
		app.script.data.WriteString(app.query)
		app.script.data.WriteString(".")
		app.script.data.WriteString(method)
		app.script.data.WriteString("();")
		return
	}
	app.script.data.WriteString(`document.getElementById("`)
	app.script.data.WriteString(app.ID())
	app.script.data.WriteString(`").`)
	app.script.data.WriteString(method)
	app.script.data.WriteString(`();`)
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

func (script *Script) Run(f string) {
	script.data.WriteString(f)
	script.data.WriteString(`();`)
}

func (script *Script) Query(q string) *App {
	sa := new(App)
	sa.query = `document.querySelector("`+q+`")`
	sa.script = script
	sa.style = &style.Style{Css: &scriptCss{script:script, app:sa}}
	
	return sa
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
	sa.Style = style.Style{Css: CssWrapper{app: sa}}
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
