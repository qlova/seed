package main

import (
	"fmt"
	"os"
	"strings"
)

func Format(property string) (upper, camel string) {
	var parts = strings.Split(property, "-")
	for i := range parts {
		upper += strings.Title(parts[i])

		if i == 0 {
			camel += parts[i]
		} else {
			camel += strings.Title(parts[i])
		}
	}
	return
}

func Exception(property string) string {
	switch property {
	case "will-change":
		return `
func (style Style) SetWillChange(properties ...interface{}) {
	var names string

	for i, property := range properties {
		var s = NewStyle()
		var catcher = propertyCatcher("")
		s.Stylable = &catcher
		
		reflect.ValueOf(property).Call([]reflect.Value{reflect.ValueOf(&s)})
		
		names += *((*string)(s.Stylable.(*propertyCatcher)))
		if i != len(properties) - 1 {
			names += ","
		}
	}
	style.set("will-change", unitType(names))
}		
`

	case "font-synthesis":
		return `
type fontSynthesisValue string
func (f fontSynthesisValue) String() string {
	return string(f)
}

func FontSynthesis(weight, style bool) fontSynthesisValue {
	if !weight && !style {
		return fontSynthesisValue("none")
	}
	var result fontSynthesisValue
	if weight {
		result += fontSynthesisValue("weight ")
	}
	if style {
		result += fontSynthesisValue("style")
	}
	return result
}

func (style Style) SetFontSynthesis(value fontSynthesisValue) {
	style.set("font-synthesis", value)
}
`
	case "grid-template-areas":
		return `
func (style Style) SetGridTemplateAreas(names []string) {
	if len(names) == 0 {
		style.set("grid-template-areas", unitType("none"))
		return
	}
	var result string
	for _, name := range names {
		result += name+" "
	}
	style.set("grid-template-areas", unitType(result))
}
`
	case "grid-template-columns":
		return `
func (style Style) SetGridTemplateColumns(values []gridTemplateValue) {
	if len(values) == 0 {
		style.set("grid-template-columns", unitType("none"))
		return
	}
	var result string
	for _, value := range values {
		result += value.String()+" "
	}
	style.set("grid-template-columns", unitType(result))
}
`

	case "grid-template-rows":
		return `
func (style Style) SetGridTemplateRows(values []gridTemplateValue) {
	if len(values) == 0 {
		style.set("grid-template-rows", unitType("none"))
		return
	}
	var result string
	for _, value := range values {
		result += value.String()+" "
	}
	style.set("grid-template-rows", unitType(result))
}
`
	case "quotes":
		return `
func (style Style) SetQuotes(quotes []string) {
	if len(quotes) == 0 {
		style.set("quotes", unitType("none"))
		return
	}
	var result string
	for _, quote := range quotes {
		result += strconv.Quote(quote)
	}
	style.set("quotes", unitType(result))
}
`
	case "transform-origin":
		return `
func (style Style) SetTransformOrigin(p positionValue, z ...unitValue) {
	if len(z) > 0 {
		style.set("transform-origin", unitType(p.String()+" "+z[0].String()))
	} else {
		style.set("transform-origin", p)
	}
}			
`
	case "transition-property":
		return `
func (style Style) SetTransitionProperty(properties ...interface{}) {
	var names string

	for _, property := range properties {
		var s = NewStyle()
		reflect.ValueOf(property).Call([]reflect.Value{reflect.ValueOf(&s)})
		
		for i := range s.Stylable.(Implementation) {
			names += i
		}
	}
	style.set("transform-property", unitType(names))
}
`
	}
	return ""
}

func SharedTypeFor(property string) string {

	switch property {
	case "animation-duration", "animation-delay", "transform-delay", "transform-duration":
		return "duration"

	case "background-clip", "box-sizing":
		return "box"

	case "background-color", "border-bottom-color", "border-color", "border-right-color",
		"border-top-color", "border-top-left-color", "border-top-right-color", "color",
		"column-rule-color", "outline-color", "text-decoration-color":
		return "color"

	case "background-image", "border-image-source", "list-style-image":
		return "image"

	case "background-position", "object-position", "perspective-origin":
		return "unitAndUnit"

	case "background-size", "border-image-width", "border-left-color":
		return "size"

	case "border-bottom-left-radius", "border-bottom-right-radius", "border-spacing", "grid-column-gap",
		"grid-row-gap", "outline-offset", "padding-bottom", "padding-left", "padding-right", "padding-top", "text-indent":
		return "unit"

	case "border-bottom-style", "border-left-style", "border-right-style", "border-style", "border-top-style",
		"column-rule-style", "outline-style":
		return "borderStyle"

	case "border-bottom-width", "border-left-width", "border-right-width", "border-top-width", "border-width", "outline-width":
		return "thickness"

	case "border-image-outset":
		return "uintOrUnit"

	case "bottom", "column-span", "flex-basis", "height", "left", "margin-bottom", "margin-left", "margin-right", "margin-top", "right", "top", "width":
		return "unitOrAuto"

	case "box-shadow", "text-shadow":
		return "shadow"

	case "break-after", "break-before":
		return "break"

	case "caret-color":
		return "colorOrAuto"

	case "counter-reset":
		return "name"

	case "flex-grow", "flex-shrink", "opacity", "tab-size":
		return "number"

	case "font-kerning":
		return "normalOrAuto"

	case "grid-auto-columns", "grid-auto-rows":
		return "gridAuto"

	case "grid-column-end", "grid-column-start", "grid-row-end", "grid-row-start":
		return "gridStop"

	case "grid-template", "grid-template-columns", "grid-template-rows":
		return "gridTemplate"

	case "letter-spacing":
		return "normalOrUnitOrAuto"

	case "max-height", "max-width", "min-height", "min-width", "perspective":
		return "unitOrNone"

	case "orphans", "white-space":
		return "uint"

	case "overflow", "overflow-x", "overflow-y":
		return "overflow"

	case "page-break-after", "page-break-before":
		return "pageBreak"

	case "width-spacing":
		return "normalOrUnit"

	case "z-index":
		return "integerOrAuto"
	}

	return ""
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("[usage] generate [properties/values]")
	}

	if os.Args[1] == "properties" {
		fmt.Println("/*This file is computer-generated*/")
		fmt.Println("package css")
		fmt.Println(`import "strconv"`)
		fmt.Println(`import "reflect"`)
		fmt.Println()

		var DoneType = make(map[string]struct{})

		//Generate shared types
		for property := range Properties {
			var TypeName = SharedTypeFor(property)
			var _, done = DoneType[TypeName]
			if TypeName != "" && !done {
				fmt.Println("type " + TypeName + "Value interface {")
				fmt.Println("\tstringable")
				fmt.Println("\t" + TypeName + "Value()")
				fmt.Println("}")
				fmt.Println("type " + TypeName + "Type string")
				fmt.Println("func (self " + TypeName + "Type) String() string { return string(self) }")
				fmt.Println("func (" + TypeName + "Type) " + TypeName + "Value() {}")
				fmt.Println()
				DoneType[TypeName] = struct{}{}
			}
		}
		for property := range Properties {

			if property[0] == '@' {
				continue
			}

			if exception := Exception(property); exception != "" {
				fmt.Println(exception)
				continue
			}

			upper, camel := Format(property)

			var TypeName = SharedTypeFor(property)

			if TypeName == "" {
				fmt.Println("type " + camel + "Value interface {")
				fmt.Println("\tstringable")
				fmt.Println("\t" + camel + "Value()")
				fmt.Println("}")
				fmt.Println("type " + camel + "Type string")
				fmt.Println("func (self " + camel + "Type) String() string { return string(self) }")
				fmt.Println("func (" + camel + "Type) " + camel + "Value() {}")
				fmt.Println()
				TypeName = camel
			}

			fmt.Println("func (style Style) " + upper + "() " + TypeName + "Value {")
			fmt.Println("\treturn " + TypeName + "Type(style.Get(\"" + property + "\"))")
			fmt.Println("}")
			fmt.Println("func (style Style) Set" + upper + "(value " + TypeName + "Value) {")

			fmt.Println("\tstyle.set(\"" + property + "\", value) ")
			fmt.Println("}")
		}
	}

	if os.Args[1] == "values" {
		fmt.Println("/*This file is computer-generated*/")
		fmt.Println("package css")
		fmt.Println()

		var DoneValue = make(map[string]struct{})
		var DoneType = make(map[string]struct{})

		for property := range Properties {
			var TypeName = SharedTypeFor(property)
			if _, done := DoneType[TypeName]; !done {
				DoneType[TypeName] = struct{}{}
			}

		}

		for property, values := range Properties {

			if property[0] == '@' {
				continue
			}

			values = append(values, []string{"unset", "initial", "inherit"}...)

			_, camel := Format(property)
			var TypeName = SharedTypeFor(property)

			//var SharedName bool
			if TypeName != "" {
				//SharedName = true
			} else {
				TypeName = camel
			}

			for _, value := range values {
				upper, camel := Format(value)

				if value == "0" {
					upper = "Zero"
					camel = "zero"
				}

				//Reserved Words.
				if camel == "default" {
					camel = "defaultValue"
				}
				if upper == "Style" {
					upper = "StyleProperty"
					camel = "styleProperty"
				}

				//TODO shorthands
				if value == "<shorthand>" {
					continue
				}

				//TODO Functions
				if strings.Contains(value, "(") {
					continue
				}

				//TODO Functions
				if strings.Contains(value, "[") {
					continue
				}

				//TODO Functions
				if strings.Contains(value, " ") {
					continue
				}

				var _, done = DoneValue[value]
				if !done && value[0] != '<' {
					fmt.Println(`const ` + upper + ` ` + camel + ` = 0;`)
					fmt.Println(`type ` + camel + ` byte`)
					fmt.Println(`func (` + camel + `) String() string { return "` + value + `" }`)
					DoneValue[value] = struct{}{}
				}

				if value[0] == '<' {
					camel = value[1 : len(value)-1]

					//There is a position css property, so the <position> type is renamed to <unitAndUnit>
					if camel == "position" {
						camel = "unitAndUnit"
					}
					//There is a filter css property, so the <filter> type is renamed to <filterMode>
					if camel == "filter" {
						camel = "filterMode"
					}

					if _, done := DoneType[camel]; !done {
						fmt.Println(`type ` + camel + `Type string`)
						fmt.Println(`func (s ` + camel + `Type) String() string { return string(s) }`)
						DoneType[camel] = struct{}{}
					}
					camel += "Type"
				}

				if _, done := DoneType[camel+" "+TypeName]; !done && TypeName+"Type" != camel {
					fmt.Println(`func (` + camel + `) ` + TypeName + `Value() {}`)
					fmt.Println()
					DoneType[camel+" "+TypeName] = struct{}{}
				}

			}
		}
	}
}
