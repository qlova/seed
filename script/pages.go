package script

import "github.com/qlova/script"
import "github.com/qlova/script/language"
import "github.com/qlova/script/language/javascript"

type Page struct {
	Seed
}

func (page Page) Goto() {
	page.Javascript(`goto("`+page.ID+`");`)
}

func (a Page) Equals(b Page) script.Bool {
	return a.Qlovascript.BoolFromLanguageType(Javascript.Bit{
		Expression: language.Statement(`("`+a.ID+`" == "`+b.ID+`")`),
	})
}