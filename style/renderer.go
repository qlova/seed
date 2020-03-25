package style

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
)

type harvester struct {
	queries map[string]*bytes.Buffer
}

func newHarvester() harvester {
	return harvester{
		make(map[string]*bytes.Buffer),
	}
}

func (h harvester) harvest(c seed.Seed) map[string]*bytes.Buffer {
	var data data
	c.Read(&data)

	for query, rules := range data.queries {
		if b, ok := h.queries[query]; !ok {
			b = new(bytes.Buffer)
			h.queries[query] = b
		}

		b := h.queries[query]
		fmt.Fprintf(b, `%v {`, css.Selector(c))
		b.WriteString(rules)
		fmt.Fprint(b, `}`)
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}

	return h.queries
}

func init() {
	css.RegisterRenderer(func(c seed.Seed) []byte {
		var harvested = newHarvester().harvest(c)
		var b bytes.Buffer

		for query, rules := range harvested {
			fmt.Fprintf(&b, `%v {`, query)
			b.Write(rules.Bytes())
			fmt.Fprint(&b, `}`)
		}

		return b.Bytes()
	})
}
