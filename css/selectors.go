package css

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
)

//Select applies a selector like :hover to the seed.
func Select(selector string, rules ...Rule) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var d data
		c.Load(&d)

		switch mode, _ := client.Seed(c); mode {
		case client.AddTo:
			panic("cannot hover on client.Seed")

		case client.Undo:
			panic("cannot hover on client.Seed")
		default:
			if d.queries == nil {
				d.queries = make(map[string]Rules)
			}
			d.queries[selector] = append(d.queries[selector], rules...)
		}

		c.Save(d)
	})
}

//Hover applies the css rules when the mouse is hovering over an element.
func Hover(rules ...Rule) seed.Option {
	return Select(":hover", rules...)
}
