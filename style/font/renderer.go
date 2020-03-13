package font

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
)

type harvester struct {
	fonts map[string]Font
}

func newHarvester() harvester {
	return harvester{
		make(map[string]Font),
	}
}

func (h harvester) harvest(s seed.Any) map[string]Font {
	var data = seeds[s.Root()]

	for _, font := range data.fonts {
		h.fonts[font.name] = font
	}

	for _, child := range s.Root().Children() {
		h.harvest(child)
	}

	return h.fonts
}

func init() {
	css.RegisterRenderer(func(s seed.Any) []byte {
		var harvested = newHarvester().harvest(s)
		var b bytes.Buffer

		for _, font := range harvested {
			fmt.Fprint(&b, `@font-face {`)
			b.Write(font.Bytes())
			fmt.Fprint(&b, `}`)
		}

		return b.Bytes()
	})
}
