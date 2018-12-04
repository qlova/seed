package seed

import "github.com/qlova/seed/style/css"

var pages []Seed

func Page() Seed {
	return NewPage()
}

func NewPage() Seed {
	seed := Col()
	
	seed.page = true
	
	seed.SetHidden()
	seed.SetWillChange(css.Property.Display)
	
	pages = append(pages, seed)
	
	return seed
}
