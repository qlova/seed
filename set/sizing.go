package set

import (
	"qlova.org/seed/use/css"
	"qlova.org/seed/use/css/units"
)

//Width sets the target width of this seed.
func Width(w units.Unit) Style {
	return css.SetWidth(css.Measure(w))
}

//MaxWidth sets the maximum width of this seed.
func MaxWidth(w units.Unit) Style {
	return css.SetMaxWidth(css.Measure(w))
}

//MinWidth sets the minumum width of this seed.
func MinWidth(w units.Unit) Style {
	return css.SetMinWidth(css.Measure(w))
}

//Height sets the target height of this seed.
func Height(w units.Unit) Style {
	return css.SetHeight(css.Measure(w))
}

//MaxHeight sets the maximum height of this seed.
func MaxHeight(w units.Unit) Style {
	return css.SetMaxHeight(css.Measure(w))
}

//MinHeight sets the minumum height of this seed.
func MinHeight(w units.Unit) Style {
	return css.SetMinHeight(css.Measure(w))
}

//Size sets the target size of this seed.
func Size(w, h units.Unit) Style {
	return css.Rules{
		css.SetWidth(css.Measure(w)),
		css.SetHeight(css.Measure(h)),
	}
}
