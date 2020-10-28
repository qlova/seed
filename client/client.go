package client

import (
	"fmt"
	"time"

	"qlova.org/seed/js"
)

//Data on how to handle client events.
type Data struct {
	id string

	On map[string]js.Script
}

var unique int

//Unique returns a unique string suitable for variable names.
func Unique() string {
	unique++
	return fmt.Sprint("unique_", unique)
}

//Open asks the client to open the specified URL.
func Open(url String) Script {
	return js.Func(`window.open`).Run(url, NewString("_blank"))
}

//Print asks the client to print the current page.
func Print() Script {
	return js.Func(`window.print`).Run()
}

//Cancel cancels the current script.
func Cancel() Script {
	return js.Script(func(q js.Ctx) {
		fmt.Fprintf(q, "throw '';")
	})
}

//After runs the given scripts after the specified duration has passed.
func After(duration time.Duration, do ...Script) Script {
	return js.Global().Run("setTimeout", NewScript(do...).GetScript(), js.NewNumber(duration.Seconds()*1000))
}

//Compound values have dependent components.
type Compound interface {
	Components() []Value
}

func flatten(value Value) []Value {
	if c, ok := value.(Compound); ok {
		return FlattenComponents(c.Components()...)
	}
	return []Value{value}
}

//FlattenComponents flattens the components to their root components.
func FlattenComponents(components ...Value) []Value {
	var flattened []Value

	for _, component := range components {
		flattened = append(flattened, flatten(component)...)
	}

	return flattened
}
