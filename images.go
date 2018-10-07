package app

func Image(path string) *Web {
	return NewImage(path)
}

func NewImage(path string) *Web {
	app := New()
	app.tag = "img"
	app.attr = "src='"+path+"'"

	RegisterAsset(path)
	return app
}


func Video(path string) *Web {
	return NewVideo(path)
}

func NewVideo(path string) *Web {
	app := New()
	app.tag = "video"
	app.attr = "src='"+path+"'"

	RegisterAsset(path)
	return app
}
