package clientside

import (
	"bytes"
	"fmt"
	"sort"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

type harvester struct {
	hooks map[Address]hook
}

func newHarvester() *harvester {
	return &harvester{
		make(map[Address]hook),
	}
}

func (h *harvester) harvest(c seed.Seed) harvester {
	var data data
	c.Load(&data)

	for _, hook := range data.hooks {
		var address, _ = hook.variable.Variable()
		var update = h.hooks[address]
		update.do = js.Append(update.do, hook.do)
		update.variable = hook.variable
		update.render.Add(hook.render.Slice()...)
		h.hooks[address] = update
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}

	return *h
}

func init() {
	client.RegisterRootRenderer(func(c seed.Seed) []byte {
		var harvested = newHarvester().harvest(c)
		var b bytes.Buffer

		b.WriteString(`
		seed.variable = {};
		seed.variable.hook = {};
		seed.variable.onchange = {};

		seed.memory = {};
		seed.memory.data = {};
		seed.memory.getItem = function(key) {
			return seed.memory.data[key];
		};
		seed.memory.setItem = function(key, value) {
			seed.memory.data[key] = value;
		};
		seed.memory.clear = function() {
			seed.memory.data = {};
		};

		c.renderALL = async (q, id) => {
		
			let l = window.q.get(id);
			if (!l) {
				if (id instanceof HTMLElement) return;

				let all = document.querySelectorAll("."+id);

				for (let child of all) {
					if (!child.classList.contains("page")) {
						await c.render(q, child);
					}
				}
				return;
			}

			if (l.tagName == "TEMPLATE") {
				return;
			}
			if (!document.contains(l)) {
				return;
			}
		
			if (l.onrender) await l.onrender();

			for (let child of l.children) {
				if (!child.classList.contains("page")) {
					await c.render(q, child);
				}
			}
		}; c.r = c.render;

		c.render = async (q, id) => {
			

			let l = q.get(id);
			if (!l) {
				if (id instanceof HTMLElement) return;

				let all = document.querySelectorAll("."+id);

				for (let child of all) {
					if (!child.classList.contains("page")) {
						await c.render(q, child);
					}
				}
				return;
			}

			if (l.tagName == "TEMPLATE") {
				return;
			}
			if (!document.contains(l)) {
				return;
			}
		
			if (l.onrender) await l.onrender();

			for (let child of l.children) {
				if (!child.classList.contains("page")) {
					await c.render(q, child);
				}
			}
		}; c.r = c.render;

		c.onrender = (q, id, exe) => {
			let l = q.get(id);
			if (!l) return;

			l.onrender = exe;
		}; c.or = c.onrender;

		seed.Scope = function(parent) {
			this.parent = parent;
			this.storage = {};
			this.setItem = function(key, value) {
				this.storage[key] = value;
			};
			this.getItem = function(key) {
				return this.storage[key];
			};
			this.get = function(id) {
				return seed.get(id);
			};
			this.refresh = function() {};
			this.data = {};

			this.getset = function(key, memory) {
				let set = this.getvar(key, memory);
				if (!set) {
					set = new Set();
					this.setvar(key, memory, set);
				}
				return set;
			}

			this.getvar = function(key, memory) {
				switch (memory) {
				case "":
					return seed.memory.getItem(key);
				case "storage":
					return JSON.parse(localStorage.getItem(key));
				case "local":
					return this.getItem(key);
				default:
					console.error("invalid memory: ", memory)
				}
			};
			
			this.setvar = function(key, memory, value) {
				let old = this.getvar(key, memory);

				switch (memory) {
				case "":
					seed.memory.setItem(key, value);
					break;
				case "storage":
					localStorage.setItem(key, JSON.stringify(value));
					break;
				case "local":
					this.setItem(key, value);
					break;
				default:
					console.error("invalid memory: ", memory)
				}

				if (JSON.stringify(old) != JSON.stringify(value)) {
					if (seed.variable.onchange[key]) seed.variable.onchange[key]();
				}

				if (seed.variable.hook[key])
					for (let id of seed.variable.hook[key]) {
						c.renderALL(this, id);
					};
			};
		}; 

		seed.Ctx = seed.Scope;

		window.scope = new seed.Scope();
		q = window.scope;

		`)

		//Deterministic render.
		keys := make([]string, 0, len(harvested.hooks))
		for i := range harvested.hooks {
			keys = append(keys, string(i))
		}
		sort.Strings(keys)

		for _, key := range keys {
			hook := harvested.hooks[Address(key)]

			address, memory := hook.variable.Variable()

			var stringAddress = client.NewString(string(address))

			switch memory {
			case ShortTermMemory:
				fmt.Fprintf(&b, `q.setvar(%v, %v, %v);`, stringAddress, client.NewString(string(memory)), hook.variable.GetDefaultValue().GetValue())
			}

			var seeds = hook.render.Slice()
			var array = make(js.NewArray, len(seeds))

			for i, seed := range seeds {
				array[i] = js.NewString(client.ID(seed))
			}

			fmt.Fprintf(&b, `seed.variable.hook[%v] = %v;`, stringAddress, array)

			if hook.do != nil {
				fmt.Fprintf(&b, `seed.variable.onchange[%v] = %v;`, stringAddress, hook.do.GetScript().GetFunction())
			}
		}

		return b.Bytes()
	})
}
