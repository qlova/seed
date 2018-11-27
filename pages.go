package seed

var pages []Seed

func Page() Seed {
	return NewPage()
}

func NewPage() Seed {
	seed := Col()
	
	seed.page = true
	
	seed.SetHidden()
	
	seed.Set("will-change", "display")
	
	pages = append(pages, seed)
	
	return seed
}
