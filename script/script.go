package script

import "bytes"
import "github.com/qlova/seed/interfaces"
import "github.com/qlova/seed/style"
import "strings"
import "fmt"

type CssWrapper struct {
	app interfaces.App
}

func (c CssWrapper) Set(a, b string) {
	c.app.GetStyle().Css.Set(a, b)
}

type App struct {
	interfaces.App
	style.Style
	
	q *script
	style *style.Style
	
	query string
}

func (app *App) SetText(text string) interfaces.App {
	data = html.EscapeString(data)
	data = strings.Replace(data, "\n", "<br>", -1)
	seed.content = []byte(data)
}

func (app *App) String() string {
	return fmt.Sprint("get(\"", app.ID(), "\")")
}

func (app *App) GetParent() interfaces.App {
	return app.q.Get(app.App.GetParent())
}

func (app *App) GetChildren() []interfaces.App {
	var children = app.App.GetChildren()
	var result = make([]interfaces.App, len(children))
	for i := range children {
		result[i] = app.q.Get(children[i])
	}
	return result
}

func (app *App) Click() {
	app.Run("click");
}

func (app *App) Run(method string) {
	if app.query != "" {
		app.q.data.WriteString(app.query)
		app.q.data.WriteString(".")
		app.q.data.WriteString(method)
		app.q.data.WriteString("();")
		return
	}
	app.q.data.WriteString(`get("`)
	app.q.data.WriteString(app.ID())
	app.q.data.WriteString(`").`)
	app.q.data.WriteString(method)
	app.q.data.WriteString(`();`)
}

func (app *App) GetStyle() *style.Style {
	return app.style
}

type Script struct {
	*script
}

type script struct {
	data bytes.Buffer
}

func (q *script) Bytes() []byte {
	return q.data.Bytes()
}

func (q *script) Write(data []byte) (int, error) {
	return q.data.Write(data)
}

func (q *script) Run(f string) {
	q.data.WriteString(f)
	q.data.WriteString(`();`)
}

func (q *script) Query(query string) *App {
	sa := new(App)
	sa.query = `document.querySelector("`+query+`")`
	sa.q = q
	sa.style = &style.Style{Css: &scriptCss{q:q, app:sa}}
	
	return sa
}

func (q *script) Get(app interfaces.App) *App {
	if app == nil {
		return nil
	}
	
	if sa, ok := app.(*App); ok {
		return sa
	}
	
	sa := new(App)
	sa.q = q
	sa.Style = style.Style{Css: CssWrapper{app: sa}}
	sa.App = app
	sa.style = &style.Style{Css: &scriptCss{q:q, app:sa}}

	return sa
}

func (q *script) LogString(msg string) {
	q.data.WriteString(`console.log("`)
	q.data.WriteString(msg)
	q.data.WriteString(`"`)
}

type scriptCss struct {
	q *script
	app *App
}

func (css *scriptCss) Set(property, value string) {
	
	if strings.Contains(property, "-") {
		splits := strings.Split(property, "-")
		
		property = splits[0] + strings.Title(splits[1])
	}
	
	css.q.data.WriteString(`get("`)
	css.q.data.WriteString(css.app.ID())
	css.q.data.WriteString(`").style.`)
	css.q.data.WriteString(property)
	css.q.data.WriteByte('=')
	css.q.data.WriteByte('"')
	css.q.data.WriteString(value)
	css.q.data.WriteByte('"')
	css.q.data.WriteByte(';')
}
