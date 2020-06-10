package feed

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
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
					if (!Array.isArray(food)) food = [food];

					let i = 1;
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
	
						let offset = i;
	
						let ctx = new c.Ctx(q);
						ctx.data = piece;
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
							await exe(ctx);
						} catch(e) {
							seed.report(e, l);
						}
	
						for (let i = 0; i < nodes; i++) {
							let child = l.children[offset+i];
							seed.render(ctx, child);
						}
						
						i++;
					}
	
					l.refreshing = false;
				});
				
			}
		}; s.feed.orf = s.feed.onrefresh;`)
	})
}
