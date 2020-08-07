package css

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

//Select applies a selector like :hover to the seed.
func Select(selector string, rules ...Rule) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var d data
		c.Read(&d)

		switch c.(type) {
		case script.Seed:
			panic("cannot hover on script.Seed")

		case script.Undo:
			panic("cannot hover on script.Seed")
		default:
			if d.queries == nil {
				d.queries = make(map[string]Rules)
			}
			d.queries[selector] = append(d.queries[selector], rules...)
		}

		c.Write(d)
	})
}

//Hover applies the css rules when the mouse is hovering over an element.
func Hover(rules ...Rule) seed.Option {
	return Select(":hover", rules...)
}
