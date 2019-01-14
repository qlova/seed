package seed

import "github.com/qlova/seed/style/css"

var pages []Seed

func NewPage() Seed {
	seed := NewColumn()
	
	seed.page = true
	
	seed.SetHidden()
	seed.SetWillChange(css.Property.Display)
	
	pages = append(pages, seed)
	
	return seed
}
