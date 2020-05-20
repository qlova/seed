package state

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/signal"
)

type refresher struct {
	seed    seed.Seed
	refresh script.Script
}

type harvester struct {
	states map[Bool]struct {
		set, unset script.Script
	}
	variables map[Value]struct {
		change script.Script
	}

	refreshers []refresher
}

func newHarvester() *harvester {
	return &harvester{
		make(map[Bool]struct {
			set, unset script.Script
		}),
		make(map[Value]struct {
			change script.Script
		}),
		nil,
	}
}

func (h *harvester) buildRefresh(c seed.Seed) script.Script {
	var data data
	c.Read(&data)

	refresh := data.onrefresh

	if data.refresh {
		refresh = script.Element(c).Run("rerender")
	}

	for _, child := range c.Children() {
		var html html.Data
		c.Read(&html)

		if html.Tag != "template" {
			refresh = refresh.Append(h.buildRefresh(child))
		}
	}

	return refresh
}

func (h *harvester) buildRefreshRoot(c seed.Seed) script.Script {
	var data data
	c.Read(&data)

	refresh := data.onrefresh

	for _, child := range c.Children() {
		refresh = refresh.Append(h.buildRefresh(child))
	}

	return refresh
}

func (h *harvester) harvest(c seed.Seed) harvester {
	var data data
	c.Read(&data)

	for state, script := range data.set {
		harvest := h.states[state]
		harvest.set = harvest.set.Append(script)
		h.states[state] = harvest
	}

	for state, script := range data.unset {
		harvest := h.states[state]
		harvest.unset = harvest.unset.Append(script)
		h.states[state] = harvest
	}

	for variable, script := range data.change {
		harvest := h.variables[variable]
		harvest.change = harvest.change.Append(script)
		h.variables[variable] = harvest
	}

	if data.refresh {
		h.refreshers = append(h.refreshers, refresher{
			seed:    c,
			refresh: h.buildRefreshRoot(c),
		})
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}

	return *h
}

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		var harvested = newHarvester().harvest(c)
		var b bytes.Buffer

		b.WriteString(`
		seed.storage = {};
		seed.storage.data = {};
		seed.storage.getItem = function(key) {
			return seed.storage.data[key];
		};
		seed.storage.setItem = function(key, value) {
			seed.storage.data[key] = value;
		}
		seed.storage.clear = function() {
			seed.storage.data = {};
		}
		
		`)

		b.WriteString(`seed.state = {};`)

		for _, refresher := range harvested.refreshers {
			js.NewCtx(&b)(func(q script.Ctx) {
				q(script.Scope(refresher.seed, q).Element() + ".rerender =async function() {")
				q(refresher.refresh)
				q("};")
			})
		}

		for state, scripts := range harvested.states {
			if state.storage == "" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"] = {`, state.key)
			fmt.Fprint(&b, `set: async function() {`)
			js.NewCtx(&b)(func(q script.Ctx) {
				q(scripts.set)
				q(signal.Emit(signal.Raw("state.set." + state.key)))
			})
			fmt.Fprint(&b, `}, unset: async function() {`)
			js.NewCtx(&b)(func(q script.Ctx) {
				q(scripts.unset)
				q(signal.Emit(signal.Raw("state.unset." + state.key)))
			})
			fmt.Fprint(&b, `}, changed: async function() {`)
			js.NewCtx(&b)(func(q script.Ctx) {
				q.If(state, func(q script.Ctx) {
					fmt.Fprintf(q, `await seed.state["%v"].set();`, state.key)
				}).Else(func(q script.Ctx) {
					fmt.Fprintf(q, `await seed.state["%v"].unset();`, state.key)
				})
			})
			fmt.Fprint(&b, `}};`)
		}

		for state := range harvested.states {
			if state.storage == "" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"].changed();`, state.key)
		}

		for variable, scripts := range harvested.variables {
			if variable.storage == "" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"] = {`, variable.key)
			fmt.Fprint(&b, `changed: async function() {`)
			js.NewCtx(&b)(scripts.change)
			fmt.Fprint(&b, `}};`)
		}

		for variable, _ := range harvested.variables {
			if variable.storage == "" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"].changed();`, variable.key)
		}

		return b.Bytes()
	})
}
