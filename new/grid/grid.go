package grid

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/new/html/div"
	"qlova.org/seed/use/css"
	"qlova.org/seed/use/css/units"
)

//New returns a new grid.
func New(options ...seed.Option) seed.Seed {
	return div.New(
		css.Set("display", "grid"),

		seed.Options(options),
	)
}

//SetSize sets the size of the grid and the size of its cells.
func SetSize(w, h int, cw, ch units.Unit) css.Rules {
	return css.Rules{
		css.Set("grid-template-columns", fmt.Sprintf("repeat(%v, %v)", w, cw)),
		css.Set("grid-template-rows", fmt.Sprintf("repeat(%v, %v)", h, ch)),
	}
}

//SetGap sets the gap between cells.
func SetGap(gap units.Unit) css.Rules {
	return css.Rules{
		css.Set("grid-gap", gap.String()),
	}
}

//Box sets the position and size of this seed in relation to a parent grid.
func Box(x, y, w, h int) css.Rules {
	return css.Rules{
		css.Set("grid-column", fmt.Sprintf("%v / %v", x+1, x+w+1)),
		css.Set("grid-row", fmt.Sprintf("%v / %v", y+1, y+h+1)),
	}
}

//Pos sets the position of this seed in relation to a parent grid.
func Pos(x, y int) css.Rules {
	return css.Rules{
		css.Set("grid-column", fmt.Sprintf("%v", x+1)),
		css.Set("grid-row", fmt.Sprintf("%v", y+1)),
	}
}
