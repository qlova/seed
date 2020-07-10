package css

import (
	"bytes"
	"fmt"
	"sort"

	"qlova.org/seed"
)

type Renderer func(root seed.Seed) []byte

var renderers []Renderer

func RegisterRenderer(r Renderer) {
	renderers = append(renderers, r)
}

func render(c seed.Seed, tracker map[string]struct{}) []byte {
	var b bytes.Buffer
	var data data
	c.Read(&data)

	if data.rules != nil && data.rules.Len() > 0 {
		fmt.Fprintf(&b, "\n%v {\n", Selector(c))
		for pair := data.rules.Oldest(); pair != nil; pair = pair.Next() {
			property, value := pair.Key, pair.Value
			fmt.Fprintf(&b, "\t%v: %v;", property, value)
		}
		fmt.Fprint(&b, "}\n")
	}

	//harvest sheets.
	if len(data.sheets) > 0 {

		//Deterministic render.
		keys := make([]string, 0, len(data.sheets))
		for i := range data.sheets {
			keys = append(keys, string(i))
		}
		sort.Strings(keys)

		for _, sheet := range keys {
			if _, ok := tracker[sheet]; !ok {
				tracker[sheet] = struct{}{}

				b.WriteString(sheet)
			}
		}
	}

	for _, child := range c.Children() {
		b.Write(render(child, tracker))
	}

	return b.Bytes()
}

//Render renders a css document from the given seed as the root element.
func Render(root seed.Seed) []byte {
	var b bytes.Buffer

	b.Write(render(root, make(map[string]struct{})))

	for _, renderer := range renderers {
		b.Write(renderer(root))
	}

	return b.Bytes()
}
