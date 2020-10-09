package html

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"

	"qlova.org/seed"
	"qlova.org/seed/client"
)

//ID returns the html ID of this seed.
func ID(root seed.Seed) string {
	return client.ID(root)
}

//Render renders the html of a seed.
func Render(c seed.Seed) []byte {
	var b bytes.Buffer
	var data Data

	c.Read(&data)

	if data.Tag != "" {
		fmt.Fprintf(&b, `<%v`, data.Tag)

		if c.Used() {
			if data.ID != nil {
				if *data.ID != "" {
					fmt.Fprintf(&b, ` id=%v`, strconv.Quote(*data.ID))
				}
			} else {
				fmt.Fprintf(&b, ` id=%v`, ID(c))
			}
		}

		if data.Attributes != nil {

			//Deterministic render.
			keys := make([]string, 0, len(data.Attributes))
			for i := range data.Attributes {
				keys = append(keys, string(i))
			}
			sort.Strings(keys)

			for _, property := range keys {
				value := data.Attributes[property]
				fmt.Fprintf(&b, " %v=%v ", property, strconv.Quote(value))
			}
		}

		if data.Classes != nil {
			fmt.Fprint(&b, ` class="`)
			for _, class := range data.Classes {
				fmt.Fprintf(&b, " %v ", class)
			}
			fmt.Fprint(&b, `" `)
		}

		_, ok := c.(client.Seed)

		if data.Style != nil || ok {
			fmt.Fprint(&b, ` style="`)

			//Deterministic render.
			keys := make([]string, 0, len(data.Style))
			for i := range data.Style {
				keys = append(keys, string(i))
			}
			sort.Strings(keys)

			for _, property := range keys {
				value := data.Style[property]
				fmt.Fprintf(&b, "%v: %v;", property, value)
			}
			if ok {
				fmt.Fprint(&b, `display: none;`)
			}
			fmt.Fprint(&b, `" `)
		}

		fmt.Fprint(&b, ">")

		b.WriteString(data.InnerHTML)
	}

	for _, child := range c.Children() {
		b.Write(Render(child))
	}

	if data.Tag != "" {
		fmt.Fprintf(&b, `</%v>`, data.Tag)
	}

	return b.Bytes()
}
