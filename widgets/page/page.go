package page

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

type Widget struct {
	seed.Seed
}

func New() Widget {
	widget := seed.New()

	widget.SetClass("page")
	
	widget.SetHidden()
	widget.SetWillChange(css.Property.Display)

	widget.SetPosition(css.Fixed)
	widget.SetTop(css.Zero)
	widget.SetLeft(css.Zero)
	widget.SetWidth(css.Number(100).Vw())
	widget.SetHeight(css.Number(100).Vh())

	return  Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface) Widget {
	var Text = New()
	parent.Root().Add(Text)
	return Text
}

func (page Widget) Script(q Script) script.Page {
	return script.Page{page.Seed.Script(q)}
}

func (seed Widget) SetPage(page Page) {
	seed.OnReady(func(q Script) {
		q.Javascript(`if (!window.localStorage.getItem("*CurrentPage")) goto("`+page.id+`");`)
	})
}
