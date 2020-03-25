package html

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

//ID returns the html ID of this seed.
func ID(root seed.Seed) string {
	return script.ID(root)
}

//Render renders the html of a seed.
func Render(c seed.Seed) []byte {
	var b bytes.Buffer
	var data data

	c.Read(&data)

	if data.tag != "" {
		fmt.Fprintf(&b, `<%v`, data.tag)

		if c.Used() {
			if data.id != nil {
				if *data.id != "" {
					fmt.Fprintf(&b, ` id=%v`, strconv.Quote(*data.id))
				}
			} else {
				fmt.Fprintf(&b, ` id=%v`, ID(c))
			}
		}

		if data.attributes != nil {
			for property, value := range data.attributes {
				fmt.Fprintf(&b, " %v=%v ", property, strconv.Quote(value))
			}
		}

		if data.classes != nil {
			fmt.Fprint(&b, ` class="`)
			for _, class := range data.classes {
				fmt.Fprintf(&b, " %v ", class)
			}
			fmt.Fprint(&b, `" `)
		}

		_, ok := c.(script.Seed)

		if data.style != nil || ok {
			fmt.Fprint(&b, ` style="`)
			for property, value := range data.style {
				fmt.Fprintf(&b, "%v: %v;", property, value)
			}
			if ok {
				fmt.Fprint(&b, `display: none;`)
			}
			fmt.Fprint(&b, `" `)
		}

		fmt.Fprint(&b, ">")

		b.WriteString(data.innerHTML)
	}

	for _, child := range c.Children() {
		b.Write(Render(child))
	}

	if data.tag != "" {
		fmt.Fprintf(&b, `</%v>`, data.tag)
	}

	return b.Bytes()
}
