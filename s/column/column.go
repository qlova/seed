package column

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/css"

	"github.com/qlova/seed/s/html/div"
)

//New returns a new row.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),

		css.SetFlexShrink(css.Number(1)),

		seed.Options(options),
	)
}
