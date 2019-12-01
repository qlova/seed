package style

import (
	"bytes"
	"strings"

	"github.com/qlova/seed/style/css"
)

//Rules is a mapping between selectors and styles.
type Rules struct {
	Standard map[string][]string
	Extras   map[string]string
}

//NewRules returns a new ruleset.
func NewRules() Rules {
	return Rules{
		make(map[string][]string),
		make(map[string]string),
	}
}

//AddStyle adds a style to the rules.
func (rules Rules) AddStyle(selector string, style Style) {
	rules.Add(selector, style.CSS())
	if len(style.selectors) > 0 {
		for suffix, style := range style.selectors {
			rules.AddExtra(selector+suffix, style.CSS())
		}
	}
}

//Add adds a new entry to the rules.
func (rules Rules) Add(selector string, style css.Stylable) {
	var properties = strings.Split(string(style.Bytes()), ";")

	for _, property := range properties {

		if len(property) == 0 {
			continue
		}

		rules.Standard[property] = append(rules.Standard[property], selector)
	}
}

//AddExtra adds a new unomptimised entry to the rules.
func (rules Rules) AddExtra(selector string, style css.Stylable) {
	rules.Extras[selector] = string(style.Bytes())
}

//Get retrieved an entry from the rules.
func (rules Rules) Get(selector string) css.Stylable {
	var result = css.NewStyle()

	for style, selectors := range rules.Standard {
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
	for property, selectors := range rules.Standard {

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

	for selector, properties := range rules.Extras {
		result.WriteString(selector)
		result.WriteByte('{')
		result.WriteString(properties)
		result.WriteByte('}')
	}

	return result.Bytes()
}
