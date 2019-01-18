package seed

func NewImage(path string) Seed {
	seed := New()
	seed.tag = "img"
	seed.attr = "src='"+path+"'"

	RegisterAsset(path)
	return seed
}

func AddImageTo(seed Seed, path string) Seed {
	var image = NewImage(path)
	seed.Add(image)
	return image
}

func NewVideo(path string) Seed {
	seed := New()
	seed.tag = "video"
	seed.attr = "src='"+path+"' playsinline preload='auto'"

	RegisterAsset(path)
	return seed
}


func NewDocument(path string) Seed {
	seed := New()
	seed.tag = "embed"
	seed.attr = "src='"+path+"'"

	if path != "" {
		RegisterAsset(path)
	}
	return seed
}
