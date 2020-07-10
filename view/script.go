package view

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`

seed.view = async function(of, name, args, fragment) {
	let first = false;
	if(!fragment) fragment = "";
	if(!of.view) of.view = {};
	if(!of.view.queue) of.view.queue = [];
	if(!of.LastView) first = true;

	if (seed.debug) {
		console.log("seed.view: ", of, name, args, fragment)
	}

	//Don't goto if we are already going to something.
	if (of.NextView != null) {
		of.view.queue.push(arguments);
		return;
	}

	//If we are going to the same View then just refresh.
	if (of.CurrentView && of.CurrentView.classList.contains(name)) {
		of.NextView = null;
		of.CurrentView.args = args;
		of.CurrentView.onviewenter();
		return;
	}

	//Find the view.
	of.NextView = null;
	let template;
	for (let child of of.children) {
		if (child.tagName == "TEMPLATE") {
			if (child.content.children[0]) {
				if (child.content.children[0].classList.contains(name)) {
					of.NextView = child.content.children[0];
					template = child;
				}
			}
		}
	}

	if (!of.NextView) {
		console.error("seed.view: invalid view ", name);
		return;
	}
	of.NextView.template = template;

	if (window.flipping) flipping.read();

	of.appendChild(of.NextView);


	let promises = [];

	of.LastView = of.CurrentView;
	of.LastName = of.CurrentName;
	of.CurrentView = of.NextView;
	of.CurrentName = name;
	of.CurrentView.args = args || {};
	
	if (of.LastView) {
		if (of.LastView.onviewexit) await of.LastView.onviewexit(of);
		if (q.setvar) q.setvar(of.id+".view."+of.LastName, "", false);
	}
	{
		if (of.CurrentView.onviewenter) await of.CurrentView.onviewenter(of);
		if (q.setvar) q.setvar(of.id+".view."+name, "", true);
	}

	if (of.view.in) {
		promises.push(of.view.in);
		of.view.in = null;
	}

	if (of.view.out) {
		promises.push(of.view.out);
		of.view.out = null;
	}

	for (let promise of promises) {
		await promise;
	}
	
	if (of.LastView) {
		if (of.LastView == of.LoadingView) {
			of.LastView.style.display = "none";
		} else {
			of.LastView.template.content.appendChild(of.LastView);
		}
	}

	if (window.flipping) flipping.flip();

	//Persistence.
	localStorage.setItem(of.id+'.CurrentView', of.CurrentView.id);
	localStorage.setItem(of.id+'.LastViewTime', Date.now());
	localStorage.setItem(of.id+'.CurrentArgs', JSON.stringify(args || {}));
	localStorage.setItem(of.id+'.CurrentFrag', fragment);

	of.animating = false;
	of.NextView = null;

	if(!of.view.queue) of.view.queue = [];
	if (of.view.queue.length > 0) {
		of.view.apply(null, of.view.queue.shift());
	}
}

seed.view.ready = async function(of, id, args) {
	of.StartingView = id;
	if (!seed.view) return;

	let saved_view = window.localStorage.getItem(of.id+'.CurrentView');
	let saved_frag = window.localStorage.getItem(of.id+'.CurrentFrag');
	let saved_args = {};
	if (window.localStorage.getItem(of.id+'.CurrentArgs') && 
		window.localStorage.getItem(of.id+'.CurrentArgs') != "undefined") {
		saved_args = JSON.parse(window.localStorage.getItem(of.id+'.CurrentArgs'));
	}

	if (saved_view && saved_frag) {
		let last_time = +window.localStorage.getItem(of.id+'.LastGotoTime');
		let hibiscus = Date.now()-last_time;

		if (hibiscus > 1000*60*10) {
			window.localStorage.removeItem(of.id+'.CurrentView');
			window.localStorage.removeItem(of.id+'.CurrentArgs');
			await seed.view(of, of.StartingView);
			return;
		}

		await seed.view(of, saved_view, saved_args, saved_frag);
	} else {
		await seed.view(of, of.StartingView, args, null);
	}
}

		`)
	})
}
