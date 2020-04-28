package signal

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
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

	for signal, script := range data.handlers {
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

		for signal, script := range harvested.signals {
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
