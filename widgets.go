package app

func NewToolBar() *Web {
	app := New()
	app.SetLayout("flex")
	app.SetSticky()
	return app
}
