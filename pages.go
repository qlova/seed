package app

func NewPage() *Web {
	app := New()
	
	app.page = true
	
	return app
}
