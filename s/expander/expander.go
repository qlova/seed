package expander

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/s/html/div"
)

//New returns a new expander that expands to fill empty space.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		css.SetFlexGrow(css.Number(1)),
	)
}
