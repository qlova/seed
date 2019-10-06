package style

import (
	"bytes"
	"strings"

	"github.com/qlova/seed/style/css"
)

//Rules is a mapping between selectors and styles.
type Rules map[string][]string

//Add adds a new entry to the rules.
func (rules Rules) Add(selector string, style css.Stylable) {
	var properties = strings.Split(string(style.Bytes()), ";")

	for _, property := range properties {

		if len(property) == 0 {
			continue
		}

		rules[property] = append(rules[property], selector)
	}
}

//Get retrieved an entry from the rules.
func (rules Rules) Get(selector string) css.Stylable {
	var result = css.NewStyle()

	for style, selectors := range rules {
		for _, selector := range selectors {
			if selector == selector {
				var parts = strings.Split(style, ":")
				result.Stylable.(css.Implementation)[parts[0]] = parts[1][:len(parts[1])-1]
			}
		}
	}

	if len(result.Stylable.(css.Implementation)) == 0 {
		return nil
	}

	return result
}

//Bytes returns the rules as CSS.
func (rules Rules) Bytes() []byte {

	//Flip the Sheet data map.
	var flipped = make(map[string][]string)
	for property, selectors := range rules {

		if len(selectors) == 0 {
			panic("Error in style.Sheet.Render()")
		}

		if len(selectors) == 1 {
			flipped[selectors[0]] = append(flipped[selectors[0]], property)
		} else {
			var selector string
			//Create multi-selector
			for i, id := range selectors {
				selector += id
				if i < len(selectors)-1 {
					selector += ","
				}
			}
			flipped[selector] = append(flipped[selector], property)
		}
	}

	var result bytes.Buffer

	for selector, properties := range flipped {

		result.WriteString(selector)
		result.WriteByte('{')
		for _, property := range properties {
			result.WriteString(property)
			result.WriteByte(';')
		}
		result.WriteByte('}')
	}

	return result.Bytes()
}
