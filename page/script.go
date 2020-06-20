package page

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

func Refresh() script.Script {
	return func(q script.Ctx) {
		q("c.r(q, seed.CurrentPage);")
	}
}

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`
seed.CurrentPage = null;
seed.NextPage = null;
seed.LastPage = null;

seed.goto = async function(id, args, url) {
	if(!url) url = "";

	//Don't goto if we are already going to something.
	if (seed.NextPage != null) {
		seed.goto.queue.push(arguments);
		return;
	}

	seed.NextPage = q.get(id);
	if (!seed.NextPage) {
		console.error("seed.goto: invalid page ", id);
		return;
	}

	var Refresh = false;
	//If we are going to the same page then return.
	if (seed.CurrentPage == seed.NextPage) {

		if (JSON.stringify(seed.CurrentPage.args) == JSON.stringify(args)) {
			seed.NextPage = null;
			return;
		}

		Refresh = true;
	}

	seed.NextPage.args = args || {};

	//Check page conditions.
	if (seed.NextPage.conditions) {
		let conditions = seed.NextPage.conditions;
		for (let condition of conditions) {
			let result = await condition.test();
			if (!result) {
				seed.NextPage = null;

				if (condition.otherwise) await condition.otherwise();
				return;
			}
		}
	}

	seed.NextPage.template = seed.NextPage.parent;

	seed.NextPage.parent.parentElement.appendChild(seed.NextPage);

	//if (window.flipping) flipping.read();

	let promises = [];
	
	seed.LastPage = seed.CurrentPage;
	seed.CurrentPage = seed.NextPage;

	if (seed.LastPage) {
		if (seed.LastPage.onpageexit) await seed.LastPage.onpageexit();
		let state = seed.state["page."+seed.LastPage.className];
		if (state && state.unset) await state.unset();
	}
	{
		if (seed.CurrentPage.onpageenter) await seed.CurrentPage.onpageenter();
		let state = seed.state["page."+seed.CurrentPage.className];
		if (state && state.set) await state.set();
	}

	if (seed.goto.in) {
		promises.push(seed.goto.in);
		seed.goto.in = null;
	}

	if (seed.goto.out) {
		promises.push(seed.goto.out);
		seed.goto.out = null;
	}

	for (let promise of promises) {
		await promise;
	}

	if (seed.LastPage && !Refresh) {
		if (seed.LastPage == seed.LoadingPage) {
			seed.LastPage.style.display = "none";
		} else {
			seed.LastPage.template.content.appendChild(seed.LastPage);
		}
	}

	//try { flipping.flip(); } catch(error) {}

	//Set title and path.
	let data = seed.NextPage.dataset;
	let path = data.path;
	if (!data.path) {
		path = "/";
	}

	/*if (args.length > 0 && path != "/") {
		for (let arg of args) {
			path += "/" + arg;
		}
	}*/

	//Persistence.
	localStorage.setItem('*CurrentPage', id);
	localStorage.setItem('*LastGotoTime', Date.now());
	localStorage.setItem('*CurrentArgs', JSON.stringify(args || {}));
	localStorage.setItem('*CurrentPath', path);

	if (!seed.goto.back && seed.production) history.pushState([id, args], data.title, path+url);
	if (!seed.goto.back && !seed.production) history.replaceState([id, args], data.title, path+url);

	seed.animating = false;
	seed.NextPage = null;

	if (seed.goto.queue.length > 0) {
		seed.goto.apply(null, seed.goto.queue.shift());
	}
}

if (seed.production) {
window.addEventListener('popstate', async function (event) {
	if (window.ActivePhotoSwipe) {
		window.ActivePhotoSwipe.close();
		return;
	}

	if (event.state == null) {
		window.history.forward();
		return;
	}

	seed.goto.back = true;
	await seed.goto.apply(null, event.state);
	seed.goto.back = false;
});
};

seed.goto.queue = [];
seed.goto.back = false;

seed.goto.ready = async function(id) {
	seed.StartingPage = id;

	if (!seed.goto) return;

	let saved_page = window.localStorage.getItem('*CurrentPage');
	let saved_path = window.localStorage.getItem('*CurrentPath');
	let saved_args = {};
	if (window.localStorage.getItem('*CurrentArgs') && 
		window.localStorage.getItem('*CurrentArgs') != "undefined") {
		saved_args = JSON.parse(window.localStorage.getItem('*CurrentArgs'));
	}

	if (window.localStorage.getItem("updating")) window.localStorage.removeItem("updating");

	//Parse the URL.
	let path = window.location.pathname;

	if (path != saved_path) {
		let templates = document.querySelectorAll('template');

		templates = Array.prototype.slice.call(templates, 0);

		templates.sort(function(x, y) {
			element1 = x.content.querySelector(".page");
			element2 = y.content.querySelector(".page");
			if (!element1) {
				if (!element2) {
					return 0;
				} else {
					return 1;
				}
			} else {
				if (!element2) {
					return 0;
				} else {
					return -1;
				}
			}

			if (element1.dataset.path.length<element2.dataset.path.length) {
			return -1;
			}
			if (element1.dataset.path.length>element2.dataset.path.length) {
			return 1;
			}
			return 0;
		});

		for (let template of templates) {
			element = template.content.querySelector(".page");
			if (element) {
				if (element.dataset.path == path) {
					await seed.goto(element.id, {});
					return;
				}
				
				//Parse path values.
				if (path.startsWith(element.dataset.path)) {
					let args = {};
					let parts = path.split('/');

					for (let i in parts) {
						if (i < 2) continue;
						args[i-1] = decodeURIComponent(parts[i]);
					}

					//Parse query string

					var query = new URLSearchParams((new URL(document.location)).searchParams);
					query.forEach(function(value, key) {
						if (value == "true") {
							args[key] = true;
						} else if (value == "false") {
							args[key] = false;
						} else if (value == "undefined") {
							args[key] = null;
						} else {
							args[key] = value;
						}
					});

					var url = path.slice(element.dataset.path.length);
					if (query.toString() != "") url += "?" + query.toString();

					await seed.goto(element.id, args, url);
					return;
				}


			}
		}
	}

	if (saved_page) {
		let last_time = +window.localStorage.getItem('*LastGotoTime');
		let hibiscus = Date.now()-last_time;

		if (hibiscus > 1000*60*10) {
			window.localStorage.removeItem('*CurrentPage');
			window.localStorage.removeItem('*CurrentArgs');
			seed.CurrentPage = seed.LoadingPage;
			await seed.goto(seed.StartingPage);
			return;
		}

		await seed.goto(saved_page, saved_args);
	} else {
		await seed.goto(seed.StartingPage);
	}
}

		`)
	})
}
