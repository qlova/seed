package seed

import "github.com/qlova/seed/style/css"

var pages []Seed

func NewPage() Seed {
	seed := NewColumn()
	
	seed.page = true
	
	seed.SetHidden()
	seed.SetWillChange(css.Property.Display)

	seed.SetPosition(css.Fixed)
	seed.SetTop(css.Zero)
	seed.SetLeft(css.Zero)
	seed.SetWidth(css.Number(100).Vw())
	seed.SetHeight(css.Number(100).Vh())
	
	pages = append(pages, seed)
	
	return seed
}
