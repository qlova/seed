package css

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
)

type Renderer func(root seed.Any) []byte

var renderers []Renderer

func RegisterRenderer(r Renderer) {
	renderers = append(renderers, r)
}

func render(seed seed.Any) []byte {
	var b bytes.Buffer
	var data = seeds[seed.Root()]

	if len(data.rules) > 0 {
		fmt.Fprintf(&b, "\n%v {\n", Selector(seed))
		for property, value := range data.rules {
			fmt.Fprintf(&b, "\t%v: %v;", property, value)
		}
		fmt.Fprint(&b, "}\n")
	}

	for _, child := range seed.Root().Children() {
		b.Write(render(child))
	}

	return b.Bytes()
}

//Render renders a css document from the given seed as the root element.
func Render(root seed.Any) []byte {
	var b bytes.Buffer

	b.Write(render(root))

	for _, renderer := range renderers {
		b.Write(renderer(root))
	}

	return b.Bytes()
}
