package script

import (
	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"
	Javascript "github.com/qlova/script/language/javascript"
)

//Page is a script interface to seed.Page.
type Page struct {
	Seed

	//Hack.
	Page interface{}
}

//Back is the JS code needed for back functionality.
const Back = `
async function back() {
	if (ActivePhotoSwipe) {
		ActivePhotoSwipe.close();
		return;
	}
	
	if (!window.goto) return;

	going_back = true;

	let onback = get(current_page).onback;
	if (onback) if (onback()) return;
	
	let noback = false;
	let last_page;
	
	if (goto_history.length == 0) {
		noback = true;
	} else {
		last_page  = goto_history.pop();
		if (last_page == null || last_page == "") {
			noback = true;
		}
	}

	//Lol
	let fallback = get(current_page).dataset.back; 
	if (fallback) {
		await goto(fallback, true);
		return;
	}


			
	let old_length = goto_history.length;
			
	await goto(last_page, true);
}
`

//Back returns to the last page on the stack. Popping the current page.
func (q Ctx) Back() {
	q.Require(Back)
	q.js.Run(`back`)
}

//Goto is the JS code for goto (page switching) support.
const Goto = `
	var animating = false;
	var going_back = false;

	var animation_complete = async function() {
		animating = false;
		
		//Process goto queue.
		let next = goto_queue.shift();
		if (next != null) {
			await goto(next);
		}
	}

	var goto_queue = [];
	var goto_history = [];

	var goto_ready = false;

	var last_page = null;
	var current_page = null;
	var next_page = null;

	var going_to = null;

	var goto_exitpromise = null;

	var goto = async function(next_page_id, private) {
		//We are still waiting for the app to load.
		if (!goto_ready) {
			return;
		}

		if (!going_to) {
			await actual_goto(next_page_id, private);
		} else {
			going_to = next_page_id;
		}
	}
	
	var actual_goto = async function(next_page_id, private) {
		going_to = null;
		//We are still waiting for the app to load.
		if (!goto_ready) {
			return;
		}

		let template = get(next_page_id+":template");
		
		if (template == null || next_page_id == loading_page || !next_page_id) {
			console.error("invalid page ", next_page_id);
			next_page_id = starting_page;
			if (next_page_id == "") return;

			template = get(starting_page+":template");
			if (template == null) {
				console.error("starting page is invalid");
				return;
			}
		}
	
		if (animating) {
			goto_queue.push(next_page_id)
			return;
		}

		if (current_page == next_page_id) return;
		if (next_page == next_page_id) return;
		next_page = next_page_id;

		if (window.flipping) flipping.read();

		for (let element of template.parentElement.childNodes) {
			if (element.classList.contains("page")) {
				if (getComputedStyle(element).display != "none") {
					var resolve = function() {
						if (element.id == loading_page) {
							set(element, "display", "none")
							return;
						}
						set(element, "animation", "")
						set(element, "z-index", "")
						get(element.id+":template").content.appendChild(element);
						going_back = false;
					};
					last_page = element.id;
					
					if (element.onpageexit) {
						try {
							await element.onpageexit();
						} catch(e) {}
						if (goto_exitpromise) {
							goto_exitpromise.then(resolve);
							goto_exitpromise = null;
							break;
						}
					}
					resolve();
				}
			}
		}

		let fallback;
		if (get(current_page)) {
			fallback = get(current_page).dataset.back;
		}
		
		
		if (last_page != null && fallback != next_page_id) {
			if (!private) goto_history.push(last_page);
		}

		template.parentElement.appendChild(template.content);

		let child = get(next_page_id);
		if (onready[child.id]) {
			await onready[child.id]();
			delete onready[child.id];
		}
		try {
			if (child.onpageenter) await child.onpageenter();
		} catch(e) {}
		current_page = next_page_id;

		//Persistence.
		window.localStorage.setItem('*CurrentPage', next_page_id);
		window.localStorage.setItem('*LastGotoTime', Date.now());
			
		next_page = null;

		//Set title and path.
		let data = get(current_page).dataset;
		if (!data.path) {
			data.path = "/";
		}
		window.history.replaceState(null, data.title, data.path);

		try { flipping.flip(); } catch(error) {}
	};
`

//Goto goes to the specified page.
func (page Page) Goto() {
	var q = page.Q
	q.Require(Goto)
	q.Require(Back)
	q.Javascript("goto('" + page.ID + "');")
	if page.Page != nil {
		q.Context.AddPage(page.ID, page.Page)
	}
}

//PrivateGoto goes to the specified page without pushing to the stack.
func (page Page) PrivateGoto() {
	var q = page.Q
	q.Require(Goto)
	q.Javascript("goto('" + page.ID + "', true);")
}

//Equals returns true if page is equal to b.
func (page Page) Equals(b Page) qlova.Bool {
	return page.Q.BoolFromLanguageType(Javascript.Bit{
		Expression: language.Statement(`("` + page.ID + `" == "` + b.ID + `")`),
	})
}

//SetCurrent sets the current page to this page. This is a low-level API and shouldn't be called. Use Goto instead.
func (page Page) SetCurrent() {
	page.Javascript(`current_page = ` + page.ID + ";")
}

//CurrentPage returns the current page.
func (q Ctx) CurrentPage() Page {
	return Page{Seed{
		ID: `"+current_page+"`,
		Q:  q,
	}, nil}
}

//ClearHistory clears the page history, you should call this after transitioning from a sign-in page.
func (q Ctx) ClearHistory() {
	q.Javascript(`goto_history = [];`)
}

//PushHistory pushes the page to history.
func (q Ctx) PushHistory(page Page) {
	q.Javascript(`goto_history.push('` + page.ID + `');`)
}

//LastPage returns the last page.
func (q Ctx) LastPage() Page {
	return Page{Seed{
		ID: `"+last_page+"`,
		Q:  q,
	}, nil}
}

//NextPage returns the next page.
func (q Ctx) NextPage() Page {
	return Page{Seed{
		ID: `"+next_page+"`,
		Q:  q,
	}, nil}
}
