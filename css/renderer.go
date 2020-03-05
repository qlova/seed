package css

import (
	"bytes"
	"fmt"

	"github.com/qlova/seed"
)

//Render renders a css document from the given seed as the root element.
func Render(seed seed.Any) []byte {
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
		b.Write(Render(child))
	}

	return b.Bytes()
}
