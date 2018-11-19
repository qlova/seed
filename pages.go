package seed

var pages []Seed

func Page() Seed {
	return NewPage()
}

func NewPage() Seed {
	seed := New()
	
	seed.page = true
	
	pages = append(pages, seed)
	
	return seed
}
