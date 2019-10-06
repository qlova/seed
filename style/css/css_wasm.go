package css

import (
	"strings"
	"syscall/js"
)

type WasmImplementation struct {
	element js.Value
}

func dashes2camels(s string) string {
	var camel string
	var parts = strings.Split(s, "-")
	for i, part := range parts {
		if i == 0 {
			camel += part
		} else {
			camel += strings.Title(part)
		}
	}
	return camel
}

func (impl WasmImplementation) Set(property, value string) {
	property = dashes2camels(property)

	impl.element.Get("style").Set(property, value)
}

func (impl WasmImplementation) Get(property string) string {
	property = dashes2camels(property)

	return js.Global().Call("getComputedStyle", impl.element).Get(property).String()
}

func (impl WasmImplementation) Bytes() []byte {
	return []byte(js.Global().Call("getComputedStyle", impl.element).Get("cssText").String())
}

func StyleOf(element js.Value) Style {
	return Style{
		Stylable: WasmImplementation{
			element: element,
		},
	}
}
