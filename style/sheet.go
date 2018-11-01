package style

import "strings"
import "bytes"

//Css Stylesheet.
type Sheet struct {
	data map[string][]string
}

func (sheet *Sheet) Render() bytes.Buffer {
	
	//Flip the Sheet data map.
	var flipped = make(map[string][]string)
	for property, selectors := range sheet.data {
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
	
	return result
}

//Main optimisation method.
func (sheet *Sheet) CreateAndReturnClassesFor(id, style string) string {
	if sheet.data == nil {
		sheet.data = make(map[string][]string)
	}
	
	var properties = strings.Split(style, ";")
	for _, property := range properties {
		
		if len(property) == 0 {
			continue
		}
		
		sheet.data[property] = append(sheet.data[property], "#"+id)
	}
	
	return ""
}
