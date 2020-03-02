package script

//Element is an HTML element.
type Element struct {
	query string
	q     Ctx
}

//Query allows finding an element based on a query string.
func (q Ctx) Query(query string) Element {
	return Element{query: query, q: q}
}

//Run calls a method on an Element.
func (element Element) Run(method string) {
	element.q.Write([]byte(`document.querySelector(` + element.query + `).` + method + `();`))
}
