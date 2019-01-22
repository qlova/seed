package seed

import "github.com/qlova/seed/script"
import "github.com/qlova/seed/style/css"

type Page struct {
	Seed
}

func NewPage() Page {
	seed := NewColumn()
	
	seed.page = true
	seed.class = "page"
	
	seed.SetHidden()
	seed.SetWillChange(css.Property.Display)

	seed.SetPosition(css.Fixed)
	seed.SetTop(css.Zero)
	seed.SetLeft(css.Zero)
	seed.SetWidth(css.Number(100).Vw())
	seed.SetHeight(css.Number(100).Vh())
	
	return Page{seed}
}

func AddPageTo(parent Interface) Page {
	var page = NewPage()
	parent.GetSeed().Add(page)
	return page
}

func (page Page) Script(q Script) script.Page {
	return script.Page{page.Seed.Script(q)}
}