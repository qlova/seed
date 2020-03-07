package align

import (
	"github.com/qlova/seed/css"
)

//Center center's the seed.
func Center() css.Rule {
	return css.SetAlignSelf(css.Center)
}
