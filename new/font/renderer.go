package font

import (
	"bytes"
	"fmt"
	"sort"

	"qlova.org/seed"
	"qlova.org/seed/use/css"
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
	c.Load(&data)

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

		//Deterministic render.
		keys := make([]string, 0, len(harvested))
		for i := range harvested {
			keys = append(keys, i)
		}
		sort.Strings(keys)

		for _, key := range keys {
			font := harvested[key]

			fmt.Fprint(&b, `@font-face {`)
			b.Write(font.Bytes())
			fmt.Fprint(&b, `}`)
		}

		return b.Bytes()
	})
}
