package app

func NewImage(path string) *Web {
	app := New()
	app.tag = "img"
	app.attr = "src='"+path+"'"
	return app
}
