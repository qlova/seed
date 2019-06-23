package style

import "github.com/qlova/seed/style/css"

import (
	"bytes"
	"strings"
)

//A stylesheet that produces optimally compressed CSS for Qlovaseed.
type Sheet map[string][]string

//Main optimisation method.
func (sheet Sheet) Add(selector string, style css.Stylable) {
	var properties = strings.Split(string(style.Bytes()), ";")

	for _, property := range properties {

		if len(property) == 0 {
			continue
		}

		sheet[property] = append(sheet[property], selector)
	}
}

func (sheet Sheet) Get(selector string) css.Stylable {
	var result = css.NewStyle()

	for style, selectors := range sheet {
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

func (sheet Sheet) Bytes() []byte {

	//Flip the Sheet data map.
	var flipped = make(map[string][]string)
	for property, selectors := range sheet {

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
