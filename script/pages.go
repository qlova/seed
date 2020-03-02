package script

import (
	"github.com/qlova/script"
	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"
)

//Page is a script interface to seed.Page.
type Page struct {
	Seed

	//Hack to allow pages to be lazy added to apps.
	Page interface{}
}

//Back is the JS code needed for back functionality.
const Back = `
seed.back = async function() {
	
	if (!seed.goto) return;

	seed.back.going = true;

	let onback = seed.CurrentPage.onback;
	if (onback) if (onback()) return;
	
	let deadend = false;
	let GotoArgs;
	
	if (seed.goto.history.length == 0) {
		deadend = true;
	} else {
		GotoArgs = seed.goto.history.pop();
		if (!GotoArgs) deadend = true;
	}

	//Lol
	let fallback = seed.CurrentPage.dataset.back; 
	if (fallback) {
		await seed.goto(fallback);
		seed.goto.history.pop();
		return;
	}

	
	if (!deadend) {
		await seed.goto.apply(null, GotoArgs);
		seed.goto.history.pop();
	}

	seed.back.going = false;
}

seed.back.going = false;
`

//Back returns to the last page on the stack. Popping the current page.
func (q Ctx) Back() {
	q.Require(Goto)
	q.Javascript(`seed.back.going = true;`)
	q.js.Run(`window.history.back`)
}

//Goto is the JS code for goto (page switching) support.
const Goto = `
	seed.CurrentPage = null;
	seed.NextPage = null;
	seed.LastPage = null;

	seed.goto = async function(id) {
		//We are still waiting for the app to load.
		if (!seed.goto.ready) {
			return;
		}

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

		let NextPageTemplate = get(id+":template");
		if (NextPageTemplate == null || id == loading_page || !id) {
			console.error("seed.goto: invalid page ", id, NextPageTemplate);

			id = starting_page;
			if (id == "") {
				console.error("seed.goto: no starting page to fallback to");
				return;
			}

			NextPageTemplate = get(starting_page+":template");
			if (NextPageTemplate == null) {
				console.error("seed.goto: starting page is invalid");
				going_to = null;
				return;
			}

			return;
		}

		NextPageTemplate.parentElement.appendChild(NextPageTemplate.content);

		seed.NextPage = get(id);

		//If we are going to the same page then return.
		if (seed.CurrentPage == seed.NextPage) {
			if (JSON.stringify(seed.CurrentPage.args) == JSON.stringify(args)) {
				seed.NextPage = null;
				return;
			}
		}

		if (window.flipping) flipping.read();

		let promises = [];

		if (onready[seed.NextPage.id]) {
			promises.push(onready[seed.NextPage.id]());
			delete onready[seed.NextPage.id];
		}
		
		seed.LastPage = seed.CurrentPage;
		seed.CurrentPage = seed.NextPage;
		seed.CurrentPage.args = args;

		if (seed.LastPage && seed.LastPage.onpageexit) await seed.LastPage.onpageexit();
		if (seed.CurrentPage.onpageenter) await seed.CurrentPage.onpageenter();

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
			if (seed.LastPage == loading_page) {
				seed.LastPage.style.display = "none";
			} else {
				get(seed.LastPage.id+":template").content.appendChild(seed.LastPage);
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

		if (!seed.goto.back && production) history.pushState([seed.CurrentPage.id].concat(seed.CurrentPage.args), data.title, path);

		seed.animating = false;
		seed.NextPage = null;

		if (seed.goto.queue.length > 0) {
			seed.goto.apply(null, seed.goto.queue.shift());
		}
	}

	if (production) {
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
	seed.goto.ready = false;
	seed.goto.back = false;
` + Back

//Arg returns the ith argument to this page.
func (page Page) Arg(i Int) String {
	return page.Q.Value("%v.args[%v]", page.Element(), i).String()
}

//Args returns the number of arguments to this page.
func (page Page) Args() Int {
	return page.Q.Value("%v.args.length", page.Element()).Int()
}

//Goto goes to the specified page.
//The provided arguments are provided to the page.
func (page Page) Goto(args ...String) {
	var q = page.Q
	q.Require(Goto)
	q.Require(Back)

	var arguments = []script.Value{q.String(page.ID)}
	for _, arg := range args {
		arguments = append(arguments, arg)
	}

	q.js.Run("await seed.goto", arguments...)

	if page.Page != nil {
		q.Context.AddPage(page.ID, page.Page)
	}
}

//Equals returns true if page is equal to b.
func (page Page) Equals(b Page) qlova.Bool {
	return script.Bool{
		language.Expression(page.Q, `(`+page.Element()+` == `+b.Element()+`)`),
	}
}

//SetCurrent sets the current page to this page. This is a low-level API and shouldn't be called. Use Goto instead.
func (page Page) SetCurrent() {
	page.Javascript(`seed.CurrentPage = ` + page.ID + ";")
}

//CurrentPage returns the current page.
func (q Ctx) CurrentPage() Page {
	return Page{Seed{
		Native: `seed.CurrentPage`,
		Q:      q,
	}, nil}
}

//ClearHistory clears the page history, you should call this after transitioning from a sign-in page.
func (q Ctx) ClearHistory() {
	q.Javascript(`history.go(-(history.length - 1));`)
}

//PushHistory pushes the page to history.
func (q Ctx) PushHistory(page Page) {
	q.Javascript(`seed.goto.history.push('` + page.ID + `');`)
}

//LastPage returns the last page.
func (q Ctx) LastPage() Page {
	return Page{Seed{
		Native: `seed.LastPage`,
		Q:      q,
	}, nil}
}

//NextPage returns the next page.
func (q Ctx) NextPage() Page {
	return Page{Seed{
		Native: `seed.NextPage`,
		Q:      q,
	}, nil}
}
