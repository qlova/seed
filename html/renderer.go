package html

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"

	"github.com/qlova/seed"
)

//ID returns the html ID of this seed.
func ID(root seed.Seed) string {
	return base64.RawURLEncoding.EncodeToString(big.NewInt(int64(root)).Bytes())
}

//Render renders the html of a seed.
func Render(seed seed.Any) []byte {
	var b bytes.Buffer
	var data = seeds[seed.Root()]

	if data.tag != "" {
		fmt.Fprintf(&b, `<%v`, data.tag)

		if seed.Root().Used() {
			if data.id != nil {
				if *data.id != "" {
					fmt.Fprintf(&b, ` id=%v`, strconv.Quote(*data.id))
				}
			} else {
				fmt.Fprintf(&b, ` id=%v`, ID(seed.Root()))
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

		if data.style != nil {
			fmt.Fprint(&b, ` style="`)
			for property, value := range data.style {
				fmt.Fprintf(&b, "%v: %v;", property, value)
			}
			fmt.Fprint(&b, `" `)
		}

		fmt.Fprint(&b, ">")

		b.WriteString(data.innerHTML)
	}

	for _, child := range seed.Root().Children() {
		b.Write(Render(child))
	}

	if data.tag != "" {
		fmt.Fprintf(&b, `</%v>`, data.tag)
	}

	return b.Bytes()
}
