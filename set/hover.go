package set

import (
	"qlova.org/seed"
	"qlova.org/seed/use/css"
)

//OnHover applies the given Styles when the user is hovering over the seed.
func OnHover(styles ...Style) seed.Option {
	var rules []css.Rule
	for i := range styles {
		rules = append(rules, styles[i].Rules()...)
	}
	return css.Hover(rules...)
}
