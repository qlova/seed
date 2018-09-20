package app

func ToolBar() *Web {
	return NewToolBar()
}

func NewToolBar() *Web {
	app := New()
	app.SetLayout("flex")
	app.SetSticky()
	return app
}


func Text() *Web {
	app := New()
	app.tag = "p"
	return app
}

func Header() *Web {
	app := New()
	app.tag = "h1"
	return app
}

func FilePicker(types string) *Web {
	app := New()
	app.tag = "input"
	app.attr = `type="file" accept="`+types+`"`
	return app
}
