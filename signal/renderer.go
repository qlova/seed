package signal

import (
	"bytes"
	"fmt"
	"sort"

	"qlova.org/seed"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

type harvester struct {
	signals map[Type]script.Script
}

func newHarvester() harvester {
	return harvester{
		make(map[Type]script.Script),
	}
}

func (h harvester) harvest(c seed.Seed) harvester {
	var data data
	c.Read(&data)

	//Deterministic render.
	keys := make([]string, 0, len(data.handlers))
	for i := range data.handlers {
		keys = append(keys, string(i.string))
	}
	sort.Strings(keys)

	for _, key := range keys {
		signal := Type{key}
		script := data.handlers[signal]
		h.signals[signal] = h.signals[signal].Append(script)
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}

	return h
}

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		var harvested = newHarvester().harvest(c)
		var b bytes.Buffer

		b.WriteString(`seed.signal = {`)

		//Deterministic render.
		keys := make([]string, 0, len(harvested.signals))
		for i := range harvested.signals {
			keys = append(keys, string(i.string))
		}
		sort.Strings(keys)

		for _, key := range keys {
			signal := Type{key}
			script := harvested.signals[signal]

			fmt.Fprintf(&b, `"%v": async function() { try{`, signal.string)
			js.NewCtx(&b)(script)
			fmt.Fprint(&b, `} catch(e) {
				seed.report(e, seed.active);
			}},`)
		}

		b.WriteString(`};`)

		return b.Bytes()
	})
}
