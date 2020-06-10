package state

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
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

	return refresh
}

func (h *harvester) buildRefreshRoot(c seed.Seed) script.Script {
	var data data
	c.Read(&data)

	refresh := data.onrefresh

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
		};
		seed.storage.clear = function() {
			seed.storage.data = {};
		};

		c.render = async (q, id) => {
			let l = q.get(id);
			if (!l) {
				if (id instanceof HTMLElement) return;

				let all = document.querySelectorAll("."+id);

				for (let child of all) await c.render(q, child);
				return;
			}
			if (l.onrender) await l.onrender();

			for (let child of l.children) await c.render(q, child);
		}; c.r = c.render;

		c.onrender = (q, id, exe) => {
			let l = q.get(id);
			if (!l) return;

			l.onrender = exe;
		}; c.or = c.onrender;
		`)

		b.WriteString(`seed.state = {};`)

		//Init the onrender function for every seed.
		for _, refresher := range harvested.refreshers {
			js.NewCtx(&b)(js.Run(js.Func("c.or"), js.NewValue("q"), js.NewString(script.ID(refresher.seed)), js.NewFunction(refresher.refresh)))
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
			fmt.Fprint(&b, `}, changed: async function(q) {`)
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
			if state.storage == "" || state.storage == "scope" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"].changed(scope);`, state.key)
		}

		for variable, scripts := range harvested.variables {
			if variable.storage == "" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"] = {`, variable.key)
			fmt.Fprint(&b, `changed: async function(q) {`)
			js.NewCtx(&b)(scripts.change)
			fmt.Fprint(&b, `}};`)
		}

		for variable := range harvested.variables {
			if variable.storage == "" || variable.storage == "scope" {
				continue
			}
			fmt.Fprintf(&b, `seed.state["%v"].changed(q);`, variable.key)
		}

		return b.Bytes()
	})
}
