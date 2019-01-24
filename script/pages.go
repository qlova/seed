package script

import qlova "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type Page struct {
	Seed
}

func (page Page) Goto() {
	page.Javascript(`goto("`+page.ID+`");`)
}

func (a Page) Equals(b Page) qlova.Bool {
	return a.Q.BoolFromLanguageType(Javascript.Bit{
		Expression: language.Statement(`("`+a.ID+`" == "`+b.ID+`")`),
	})
}

func (page Page) SetCurrent() {
	page.Javascript(`current_page = `+page.ID)
}