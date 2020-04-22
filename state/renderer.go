package state

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
)

type refresher struct {
	seed    seed.Seed
	refresh script.Script
}

type harvester struct {
	states map[State]struct {
		set, unset script.Script
	}
	variables map[Value]struct {
		change script.Script
	}

	refreshers []refresher
}

func newHarvester() *harvester {
	return &harvester{
		make(map[State]struct {
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
			refresh: h.buildRefresh(c),
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

		b.WriteString(`seed.state = {};`)

		for state, scripts := range harvested.states {
			if state.storage == "" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"] = {`, state.key)
			fmt.Fprint(&b, `set: async function() {`)
			js.NewCtx(&b)(scripts.set)
			fmt.Fprint(&b, `}, unset: async function() {`)
			js.NewCtx(&b)(scripts.unset)
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

		for _, refresher := range harvested.refreshers {
			js.NewCtx(&b)(func(q script.Ctx) {
				q(script.Scope(refresher.seed, q).Element() + ".rerender = function() {")
				q(refresher.refresh)
				q("};")
			})
		}

		return b.Bytes()
	})
}
