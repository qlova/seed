package app

func Page() *Web {
	return NewPage()
}

func NewPage() *Web {
	app := New()
	
	app.page = true
	
	return app
}
