package script

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type Page struct {
	Seed
}

func (page Page) Goto() {
	page.Javascript(`goto("` + page.ID + `");`)
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
