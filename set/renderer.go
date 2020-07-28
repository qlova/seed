package set

import (
	"bytes"
	"fmt"
	"sort"

	"qlova.org/seed"
	"qlova.org/seed/css"
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

	//Deterministic render.
	keys := make([]string, 0, len(data.queries))
	for i := range data.queries {
		keys = append(keys, string(i))
	}
	sort.Strings(keys)

	for _, query := range keys {
		rules := data.queries[query]

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

		//Deterministic render.
		keys := make([]string, 0, len(harvested))
		for i := range harvested {
			keys = append(keys, string(i))
		}
		sort.Strings(keys)

		for _, query := range keys {
			rules := harvested[query]

			fmt.Fprintf(&b, `%v {`, query)
			b.Write(rules.Bytes())
			fmt.Fprint(&b, `}`)
		}

		return b.Bytes()
	})
}
