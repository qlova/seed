package set

import (
	"fmt"

	"qlova.org/seed/web/css"
	"qlova.org/seed/web/css/units"
)

//Translation sets the translation of this seed.
func Translation(x, y units.Unit) css.Rules {
	return css.Rules{
		css.Set("--x", css.Measure(x).String()),
		css.Set("--y", css.Measure(y).String()),
		css.Set("transform", "translate(var(--x, 0), var(--y, 0)) rotate(var(--angle, 0)) scale(var(--scale, 1), var(--scale, 1))"),
	}
}

//Scale sets the scale of this seed.
func Scale(factor float64) css.Rules {
	return css.Rules{
		css.Set("--scale", fmt.Sprint(factor)),
		css.Set("transform", "translate(var(--x, 0), var(--y, 0)) rotate(var(--angle, 0)) scale(var(--scale, 1), var(--scale, 1))"),
	}
}

//Angle sets the rotation angle of this seed.
func Angle(angle float64) css.Rules {
	return css.Rules{
		css.Set("--angle", fmt.Sprintf(`%frad`, angle)),
		css.Set("transform", "translate(var(--x, 0), var(--y, 0)) rotate(var(--angle, 0)) scale(var(--scale, 1), var(--scale, 1))"),
	}
}
