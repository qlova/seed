package app

func NewImage(path string) *App {
	app := New()
	app.tag = "img"
	app.attr = "src='"+path+"'"
	return app
}
