package css

import "strings"
import "github.com/gopherjs/gopherjs/js"

type JSImplementation struct {
	element js.Object
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

func (impl JSImplementation) Set(property, value string) {
	property = dashes2camels(property)

	impl.element.Get("style").Set(property, value)
}

func (impl JSImplementation) Get(property string) string {
	property = dashes2camels(property)

	return js.Global.Call("getComputedStyle", impl.element).Get(property).String()
}

func (impl JSImplementation) Bytes() []byte {
	return []byte(js.Global.Call("getComputedStyle", impl.element).Get("cssText").String())
}

func StyleOf(element js.Object) Style {
	return Style{
		Stylable: JSImplementation{
			element: element,
		},
	}
}
