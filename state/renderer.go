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

func (h harvester) harvest(s seed.Any) harvester {
	var data = seeds[s.Root()]

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

	for _, child := range s.Root().Children() {
		h.harvest(child)
	}

	return h
}

func init() {
	script.RegisterRenderer(func(s seed.Any) []byte {
		var harvested = newHarvester().harvest(s)
		var b bytes.Buffer

		b.WriteString(`seed.state = {};`)

		for state, scripts := range harvested.states {
			fmt.Fprintf(&b, `seed.state["%v"] = {`, state.key)
			fmt.Fprint(&b, `set: async function() {`)
			b.Write(script.ToJavascript(scripts.set))
			fmt.Fprint(&b, `}, unset: async function() {`)
			b.Write(script.ToJavascript(scripts.unset))
			fmt.Fprint(&b, `}, changed: async function() {}};`)
		}

		for state := range harvested.states {
			b.Write(script.ToJavascript(func(q script.Ctx) {
				q.If(state, func() {
					state.Set(q)
				}).Else(func() {
					state.Unset(q)
				}).End()
			}))
		}

		for variable, scripts := range harvested.variables {
			fmt.Fprintf(&b, `seed.state["%v"] = {`, variable.key)
			fmt.Fprint(&b, `changed: async function() {`)
			b.Write(script.ToJavascript(scripts.change))
			fmt.Fprint(&b, `}};`)
		}

		for variable, _ := range harvested.variables {
			fmt.Fprintf(&b, `seed.state["%v"].changed();`, variable.key)
		}

		return b.Bytes()
	})
}
