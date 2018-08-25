package app

import "bytes"

type ScriptApp struct {
	script *Script

	Web
}

func (app *ScriptApp) GetParent() App {
	return app.script.Get(app.Web.GetParent())
}

func (app *ScriptApp) GetChildren() []App {
	var children = app.Web.GetChildren()
	var result = make([]App, len(children))
	for i := range children {
		result[i] = app.script.Get(children[i])
	}
	return result
}


type Script struct {
	data bytes.Buffer
}

func (script *Script) Bytes() []byte {
	return script.data.Bytes()
}

func (script *Script) Get(app App) *ScriptApp {
	if app == nil {
		return nil
	}
	
	if sa, ok := app.(*ScriptApp); ok {
		return sa
	}
	
	sa := new(ScriptApp)
	sa.script = script
	
	sa.Web = Web{}
	
	sa.Web.id = app.ID()
	sa.Web.Style.css = &ScriptCss{script:script, app:sa}
	
	sa.Web.parent = app.GetParent()
	sa.Web.children = app.GetChildren()
	sa.Web.page = app.Page()

	return sa
}

func (script *Script) LogString(msg string) {
	script.data.WriteString(`console.log("`)
	script.data.WriteString(msg)
	script.data.WriteString(`"`)
}
