package script

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/js"
)

//Element returns the js Element of a seed.
func Element(c seed.Seed) js.Element {
	return js.Element{Value: js.Call(`seed.get`, js.NewString(ID(c)))}
}
