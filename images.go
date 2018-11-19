package seed

func Image(path string) Seed {
	return NewImage(path)
}

func NewImage(path string) Seed {
	seed := New()
	seed.tag = "img"
	seed.attr = "src='"+path+"'"

	RegisterAsset(path)
	return seed
}


func Video(path string) Seed {
	return NewVideo(path)
}

func NewVideo(path string) Seed {
	seed := New()
	seed.tag = "video"
	seed.attr = "src='"+path+"' playsinline preload='auto'"

	RegisterAsset(path)
	return seed
}
