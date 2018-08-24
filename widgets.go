package app

func NewToolBar() *App {
	app := New()
	app.SetLayout("flex")
	app.SetSticky()
	return app
}
