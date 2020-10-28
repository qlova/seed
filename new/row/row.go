package row

import (
	"qlova.org/seed"
	"qlova.org/seed/use/css"

	"qlova.org/seed/new/html/div"
)

//New returns a new row.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Row),

		css.SetFlexShrink(css.Number(1)),

		seed.Options(options),
	)
}

//Set returns an option that sets the seed to layout children in a row.
func Set() css.Rule {
	return css.SetFlexDirection(css.Row)
}

//Wrap returns an option that sets the seed to wrap it's children into multiple rows.
func Wrap() css.Rule {
	return css.SetFlexWrap(css.Wrap)
}
