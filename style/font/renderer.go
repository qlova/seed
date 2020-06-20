package font

import (
	"bytes"
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/css"
)

type harvester struct {
	fonts map[string]Font
}

func newHarvester() harvester {
	return harvester{
		make(map[string]Font),
	}
}

func (h harvester) harvest(c seed.Seed) map[string]Font {
	var data data
	c.Read(&data)

	for _, font := range data.fonts {
		h.fonts[font.name] = font
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}

	return h.fonts
}

func init() {
	css.RegisterRenderer(func(c seed.Seed) []byte {
		var harvested = newHarvester().harvest(c)
		var b bytes.Buffer

		for _, font := range harvested {
			fmt.Fprint(&b, `@font-face {`)
			b.Write(font.Bytes())
			fmt.Fprint(&b, `}`)
		}

		return b.Bytes()
	})
}
