package expander

import (
	"qlova.org/seed"
	"qlova.org/seed/web/css"
	"qlova.org/seed/new/html/div"
)

//New returns a new expander that expands to fill empty space.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		css.SetFlexGrow(css.Number(1)),

		seed.Options(options),
	)
}

//Set sets the seed to expand to fill up space.
func Set() css.Rule {
	return css.SetFlexGrow(css.Number(1))
}

//SetRatio sets the ratio of remaining space the expander should fill up.
func SetRatio(ratio float64) css.Rule {
	return css.SetFlexGrow(css.Number(ratio))
}
