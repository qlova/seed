package page

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`
seed.CurrentPage = null;
seed.NextPage = null;
seed.LastPage = null;

seed.goto = async function(id) {
	//Don't goto if we are already going to something.
	if (seed.NextPage != null) {
		seed.goto.queue.push(arguments);
		return;
	}

	//Collect arguments.
	let args = [];
	if (arguments.length > 1) {
		for (let i = 1; i < arguments.length; i++) {
			args.push(arguments[i]);
		}
	}

	seed.NextPage = seed.get(id);
	if (!seed.NextPage) {
		console.error("seed.goto: invalid page ", id);
		return;
	}
	seed.NextPage.template = seed.NextPage.parent;

	seed.NextPage.parent.parentElement.appendChild(seed.NextPage);

	//If we are going to the same page then return.
	if (seed.CurrentPage == seed.NextPage) {
		if (JSON.stringify(seed.CurrentPage.args) == JSON.stringify(args)) {
			seed.NextPage = null;
			return;
		}
	}

	if (window.flipping) flipping.read();

	let promises = [];
	
	seed.LastPage = seed.CurrentPage;
	seed.CurrentPage = seed.NextPage;
	seed.CurrentPage.args = args;

	if (seed.LastPage) {
		if (seed.LastPage.onpageexit) await seed.LastPage.onpageexit();
		let state = seed.state["page."+seed.LastPage.id];
		if (state && state.unset) await state.unset();
	}
	{
		if (seed.CurrentPage.onpageenter) await seed.CurrentPage.onpageenter();
		let state = seed.state["page."+seed.CurrentPage.id];
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

	try { flipping.flip(); } catch(error) {}

	for (let promise of promises) {
		await promise;
	}

	if (seed.LastPage) {
		if (seed.LastPage == seed.LoadingPage) {
			seed.LastPage.style.display = "none";
		} else {
			seed.LastPage.template.content.appendChild(seed.LastPage);
		}
	}

	//Set title and path.
	let data = seed.NextPage.dataset;
	let path = data.path;
	if (!data.path) {
		path = "/";
	}

	if (args.length > 0 && path != "/") {
		for (let arg of args) {
			path += "/" + arg;
		}
	}

	//Persistence.
	localStorage.setItem('*CurrentPage', seed.NextPage.id);
	localStorage.setItem('*LastGotoTime', Date.now());
	localStorage.setItem('*CurrentPath', path);

	if (!seed.goto.back && seed.production) history.pushState([seed.CurrentPage.id].concat(seed.CurrentPage.args), data.title, path);
	if (!seed.goto.back && !seed.production) history.replaceState([seed.CurrentPage.id].concat(seed.CurrentPage.args), data.title, path);

	seed.animating = false;
	seed.NextPage = null;

	if (seed.goto.queue.length > 0) {
		seed.goto.apply(null, seed.goto.queue.shift());
	}
}

if (seed.production) {
window.addEventListener('popstate', async function (event) {
	if (ActivePhotoSwipe) {
		ActivePhotoSwipe.close();
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
	if (window.localStorage) {
		if (window.localStorage.getItem("updating")) {
			window.localStorage.removeItem("updating");
		}

		if (!seed.goto) return;
		let saved_page = window.localStorage.getItem('*CurrentPage');
		let saved_path = window.localStorage.getItem('*CurrentPath');
		if (saved_page && saved_path) {
			let last_time = +window.localStorage.getItem('*LastGotoTime');
			let hibiscus = Date.now()-last_time;

			if (hibiscus > 1000*60*10) {
				window.localStorage.removeItem('*CurrentPage');
				seed.CurrentPage = seed.LoadingPage;
				await seed.goto(seed.StartingPage);
				return;
			}

			let splits = saved_path.split("/");
			if (splits.length > 2) {
				await seed.goto.apply(null, [window.localStorage.getItem('*CurrentPage')].concat(window.localStorage.getItem('*CurrentPath').split("/").slice(2)));
			} else {
				await seed.goto(saved_page);
			}

			//clear history
			last_page = null;
			goto_history = [];

			if (seed.get(saved_page) && seed.get(saved_page).enterpage)
				seed.get(saved_page).enterpage();
		} else {
			await seed.goto(seed.StartingPage);
		}
	} else {
		await seed.goto(seed.StartingPage);
	}
}

		`)
	})
}
