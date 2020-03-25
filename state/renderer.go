package state

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

type harvester struct {
	states map[State]struct {
		set, unset script.Script
	}
	variables map[Value]struct {
		change script.Script
	}
}

func newHarvester() harvester {
	return harvester{
		make(map[State]struct {
			set, unset script.Script
		}),
		make(map[Value]struct {
			change script.Script
		}),
	}
}

func (h harvester) harvest(c seed.Seed) harvester {
	var data data
	c.Read(&data)

	for state, script := range data.set {
		harvest := h.states[state]
		harvest.set = harvest.set.Then(script)
		h.states[state] = harvest
	}

	for state, script := range data.unset {
		harvest := h.states[state]
		harvest.unset = harvest.unset.Then(script)
		h.states[state] = harvest
	}

	for variable, script := range data.change {
		harvest := h.variables[variable]
		harvest.change = harvest.change.Then(script)
		h.variables[variable] = harvest
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

		b.WriteString(`seed.state = {};`)

		for state, scripts := range harvested.states {
			if state.storage == "" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"] = {`, state.key)
			fmt.Fprint(&b, `set: async function() {`)
			b.Write(script.ToJavascript(scripts.set))
			fmt.Fprint(&b, `}, unset: async function() {`)
			b.Write(script.ToJavascript(scripts.unset))
			fmt.Fprint(&b, `}, changed: async function() {`)
			b.Write(script.ToJavascript(func(q script.Ctx) {
				q.If(state, func() {
					fmt.Fprintf(q, `seed.state["%v"].set();`, state.key)
				}).Else(func() {
					fmt.Fprintf(q, `seed.state["%v"].unset();`, state.key)
				}).End()
			}))
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
			b.Write(script.ToJavascript(scripts.change))
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
