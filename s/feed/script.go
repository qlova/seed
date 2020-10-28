package feed

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
)

func init() {
	client.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`s.feed = {}; s.feed.onrefresh = (q, id, template_id, feed, exe, mem, adr) => {
			let l = q.get(id);
			if (!l) return;

			let template = document.getElementById(template_id);

			if (!template) console.error("template missing", template_id);


			l.onrefresh = () => {

				//don't refresh if we are already refreshing.
				if (l.refreshing) return;
				l.refreshing = true;

				q.setvar(mem, adr, false);

				//remove previous content.
				l.textContent = "";

				feed().catch((e) => {
					seed.report(e, l);

					l.refreshing = false;
				}).then(async(food) => {
					if (!food) {
						l.refreshing = false;
						return;
					}
					if (!Array.isArray(food)) food = [food];

					if (food.length == 1) {
						if (!food[0]) {
							l.refreshing = false;
							return;
						}
					}

					if (food.length > 0) {
						await q.setvar(mem, adr, true);
					}

					let i = 0;
					for (let piece of food) {
						let clone = template.content.cloneNode(true);
						let nodes = clone.children.length;
	
						//hacky tween fix
						let update_tween = function(element) {
							let key = element.getAttribute("data-flip-key");
							if (key) element.setAttribute("data-flip-key", key+" "+i);
							for (let child of element.children) {
								update_tween(child);
							};
						};
						for (let child of clone.children) {
							update_tween(child);
						}
	
						l.appendChild(clone);
	
						let offset = i*nodes;
	
						let ctx = new c.Ctx(q);
						ctx.data = piece;
						ctx.i = offset;
						ctx.feed = food;

						ctx.nodes = [];
						for (let i = 0; i < nodes; i++) {
							let child = l.children[offset+i];
							ctx.nodes.push(child);
						}

						ctx.get = function(id) {
							if (id instanceof HTMLElement) return id;
							
							let result;
	
							for (let i = 0; i < ctx.nodes.length; i++) {
								let child = ctx.nodes[i];
								if (!child) debugger;
								if (child.className == id) return child;
								result= child.querySelector("." + id);
								if (result) return result;
							}

							let old = seed.get.cache;
							seed.get.cache = null;
							result = ctx.parent.get(id);
							seed.get.cache = old;
	
							return result;
						};
	
						try {
							let f = await exe();
							await f(ctx);
						} catch(e) {
							l.refreshing = false;
							seed.report(e, l);
						}
	
						for (let j = 0; j < nodes; j++) {
							let child = l.children[offset+j];
							child.setAttribute('data-id', i);
						}
						
						i++;
					}

					await seed.render(q, l);
					l.refreshing = false;
				});
				
			}
		}; s.feed.orf = s.feed.onrefresh;`)
	})
}
