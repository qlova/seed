package clientside

import (
	"bytes"
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
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
	c.Read(&data)

	for _, hook := range data.hooks {
		var address, _ = hook.variable.Variable()
		var update = h.hooks[address]
		update.do = update.do.Append(hook.do)
		update.variable = hook.variable
		update.render.Add(hook.render.Slice()...)
		h.hooks[address] = hook
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

			this.getvar = function(key, memory) {
				switch (memory) {
				case "":
					return seed.memory.getItem(key);
				case "storage":
					return JSON.parse(localStorage.getItem(key));
				default:
					console.error("invalid memory: ", memory)
				}
			};
			
			this.setvar = function(key, memory, value) {
				switch (memory) {
				case "":
					seed.memory.setItem(key, value);
					break;
				case "storage":
					localStorage.setItem(key, JSON.stringify(value));
					break;
				default:
					console.error("invalid memory: ", memory)
				}

				if (seed.variable.hook[key])
					for (let id of seed.variable.hook[key]) {
						c.render(this, this.get(id));
					};
				if (seed.variable.onchange[key]) seed.variable.onchange[key]();
			};
		}; 

		seed.Ctx = seed.Scope;

		window.scope = new seed.Scope();
		q = window.scope;

		`)

		for _, hook := range harvested.hooks {
			address, memory := hook.variable.Variable()

			var stringAddress = client.NewString(string(address))

			switch memory {
			case ShortTermMemory:
				fmt.Fprintf(&b, `q.setvar(%v, %v, %v);`, stringAddress, client.NewString(string(memory)), hook.variable.GetDefaultValue().GetValue())
			}

			var seeds = hook.render.Slice()
			var array = make(js.NewArray, len(seeds))

			for i, seed := range seeds {
				array[i] = js.NewString(script.ID(seed))
			}

			fmt.Fprintf(&b, `seed.variable.hook[%v] = %v;`, stringAddress, array)

			if hook.do != nil {
				fmt.Fprintf(&b, `seed.variable.onchange[%v] = %v;`, stringAddress, hook.do.GetFunction())
			}
		}

		return b.Bytes()
	})
}
