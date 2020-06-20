package column

import (
	"qlova.org/seed"
	"qlova.org/seed/css"

	"qlova.org/seed/s/html/div"
)

//New returns a new row.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),

		css.SetFlexShrink(css.Number(1)),
		//style.SetMinHeight(0),

		seed.Options(options),
	)
}
