package row

import (
	"qlova.org/seed"
	"qlova.org/seed/css"

	"qlova.org/seed/s/html/div"
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
