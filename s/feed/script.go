package feed

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`s.feed = {}; s.feed.onrefresh = (q, id, feed, exe) => {
			let l = q.get(id);
			if (!l) return;

			let template = l.children[0];

			template.onrefresh = () => {

				//don't refresh if we are already refreshing.
				if (l.refreshing) return;
				l.refreshing = true;

				//remove previous content.
				while (l.childNodes.length > 1) l.removeChild(l.lastChild);

				feed().catch((e) => {
					seed.report(e, l);

					l.refreshing = false;
				}).then(async(food) => {
					if (!food) return;
					if (!Array.isArray(food)) food = [food];

					if (food.length == 1) {
						if (!food[0]) return;
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
	
						let offset = i*nodes+1;
	
						let ctx = new c.Ctx(q);
						ctx.data = piece;
						ctx.i = offset;
						ctx.get = function(id) {
							if (id instanceof HTMLElement) return id;
							
							let result;
	
							for (let i = 0; i < nodes; i++) {
								let child = l.children[offset+i];
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
							seed.report(e, l);
						}
	
						for (let i = 0; i < nodes; i++) {
							let child = l.children[offset+i];
							await seed.render(ctx, child);
						}
						
						i++;
					}
	
					l.refreshing = false;
				});
				
			}
		}; s.feed.orf = s.feed.onrefresh;`)
	})
}
