package column

import (
	"qlova.org/seed"
	"qlova.org/seed/use/css"

	"qlova.org/seed/new/html/div"
)

//New returns a new row.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),

		css.SetFlexShrink(css.Number(1)),
		//style.SetMinHeight(0),

		css.Set("flex-basis", "auto"),
		seed.Options(options),
	)
}

//Set returns an option that sets the seed to layout children in a column.
func Set() css.Rule {
	return css.SetFlexDirection(css.Column)
}

//Wrap returns an option that sets the seed to wrap it's children into multiple columns.
func Wrap() css.Rule {
	return css.SetFlexWrap(css.Wrap)
}
