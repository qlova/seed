package toolbar

import "github.com/qlova/seed"
import "github.com/qlova/seed/style/css"

type Widget struct {
	seed.Seed
}

func New() Widget {
	var widget = seed.New()

	widget.Stylable.Set("display", "flex")
	widget.Stylable.Set("position", "fixed")
	widget.SetFlexDirection(css.Row)
	widget.SetBottom(css.Zero)
	widget.SetLeft(css.Zero)

	widget.SetWidth(100)
	widget.SetHeight(2 * seed.Em)

	widget.Set("justify-content", "space-evenly")

	return Widget{widget}
}

func AddTo(parent seed.Interface) Widget {
	var widget = New()
	widget.AddTo(parent)
	return widget
}
