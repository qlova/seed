package script

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type Page struct {
	Seed
}

const Back = `
function back() {
	if (ActivePhotoSwipe) {
		ActivePhotoSwipe.close();
		return;
	}
	
	if (!window.goto) return;
	
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
	if (noback) {
		if (fallback) goto(fallback);
		return; 
	}  

			
	let old_length = goto_history.length;
			
	goto(last_page);
	if (goto_history.length  > old_length)
	goto_history.pop();
}
`

func (q Script) Back() {
	q.Require(Back)
	q.js.Run(`back`)
}

const Goto = `
	var animating = false;

	var animation_complete = function() {
		animating = false;
		
		//Process goto queue.
		let next = goto_queue.shift();
		if (next != null) {
			goto(next);
		}
	}

	var goto_queue = [];
	var goto_history = [];

	var goto_ready = false;

	var last_page = null;
	var current_page = null;
	var next_page = null;
	var goto = function(next_page_id) {
		if (!goto_ready) {
			return;
		}
		
		if (get(next_page_id) == null || get(next_page_id).className != "page" || next_page_id == loading_page) {
			next_page_id = starting_page;
			if (next_page_id == "") return;
		}
	
		if (animating) {
			goto_queue.push(next_page_id)
			return;
		}
		if (current_page == next_page_id) return;
		if (next_page == next_page_id) return;
		next_page = next_page_id;

		for (let element of get(next_page_id).parentElement.childNodes) {
			if (element.classList.contains("page")) {
				if (getComputedStyle(element).display != "none") {
					element.style.display ='none';						
					if (element.exitpage) element.exitpage();
					last_page = element.id;
				}
			}
		}

		let fallback;
		if (get(current_page)) {
			fallback = get(current_page).dataset.back;
		}
		
		
		if (last_page != null && fallback != next_page_id) {
			goto_history.push(last_page);
		}
		
		let next_element = get(next_page_id);
		if (next_element) {
			next_element.style.display = 'inline-flex';
			if (next_element.enterpage) next_element.enterpage();
			current_page = next_page_id;

			//Persistence.
			window.localStorage.setItem('*CurrentPage', next_page_id);
			
		}
		next_page = null;
	};
`

func (page Page) Goto() {
	var q = page.Q
	q.Require(Goto)
	q.js.Run("goto", q.String(page.ID))
}

func (a Page) Equals(b Page) qlova.Bool {
	return a.Q.BoolFromLanguageType(Javascript.Bit{
		Expression: language.Statement(`("` + a.ID + `" == "` + b.ID + `")`),
	})
}

func (page Page) SetCurrent() {
	page.Javascript(`current_page = ` + page.ID + ";")
}

func (q Script) CurrentPage() Page {
	return Page{Seed{
		ID: `"+current_page+"`,
		Q:  q,
	}}
}

//Clear the page history, you should call this after transitioning from a sign-in page.
func (q Script) ClearHistory() {
	q.Javascript(`goto_history = [];`)
}
