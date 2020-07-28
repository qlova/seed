package client

import (
	"qlova.org/seed"
	"qlova.org/seed/js"
)

//Open asks the client to open the specified URL.
func Open(url String) Script {
	return js.Func(`window.open`).Run(url, NewString("_blank"))
}

//Print asks the client to print the current page.
func Print() Script {
	return js.Func(`window.print`).Run()
}

//OnRender is called whenever this seed is asked to render itself.
func OnRender(do Script) seed.Option {
	return On("render", do)
}

//Compound values have dependent components.
type Compound interface {
	Components() []Value
}

func flatten(value Value) []Value {
	if c, ok := value.(Compound); ok {
		return FlattenComponents(c.Components())
	}
	return []Value{value}
}

//FlattenComponents flattens the components to their root components.
func FlattenComponents(components []Value) []Value {
	var flattened []Value

	for _, component := range components {
		flattened = append(flattened, flatten(component)...)
	}

	return flattened
}
