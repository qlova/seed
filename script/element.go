package script

import (
	"qlova.org/seed"
	"qlova.org/seed/js"
)

//Element returns the js Element of a seed.
func Element(c seed.Seed) js.Element {
	return js.Element{Value: js.Call(js.Function{js.NewValue(`q.get`)}, js.NewString(ID(c)))}
}

//Click simulates a click of this seed.
func Click(c seed.Seed) Script {
	return Element(c).Run("click")
}
