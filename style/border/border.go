package border

import "github.com/qlova/seed/css"

//Remove removes the borders of this seed.
func Remove() css.Rule {
	return css.Set("border", "none")
}
