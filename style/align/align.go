package align

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
)

//Center center's the seed.
func Center() seed.Option {
	return css.Set("align-self", "center")
}
