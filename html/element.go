package html

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/js"
)

//Element returns the js Element of a seed.
func Element(c seed.Seed) js.Element {
	return js.Element{Value: js.Call(js.Function{js.NewValue(`q.get`)}, js.NewString(client.ID(c)))}
}
