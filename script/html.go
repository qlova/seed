package script

import (
	qlova "github.com/qlova/script"
	"github.com/qlova/script/language"
)

//Element is an HTML element.
type Element struct {
	query string
	q     Ctx
}

//Query allows finding an element based on a query string.
func (q Ctx) Query(query qlova.String) Element {
	return Element{query: raw(query), q: q}
}

//Run calls a method on an Element.
func (element Element) Run(method string) {
	element.q.Raw("Javascript", language.Statement(`document.querySelector(`+element.query+`).`+method+`();`))
}
