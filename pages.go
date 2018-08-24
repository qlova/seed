package app

func NewPage() *App {
	app := New()
	
	app.page = true
	
	return app
}
