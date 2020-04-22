package feed

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`seeds.feed = {};

		seeds.feed.refresh = async function(element, response, scripts) {
			if (!Array.isArray(response)) response = [response];
			for (let value of response) {
				let data = value;
				let clone = element.content.cloneNode(true);
				let old = seed.get;
				let parent = element.parentNode;
				seed.get = function(id) {
					let get = clone.querySelector("." + id);
					if (!get)
						return old(id);
					return get;
				};
				await scripts(data);
				seed.get = old;
				parent.appendChild(clone);
			}
		};`)
	})
}
