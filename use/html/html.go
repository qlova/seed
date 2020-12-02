package html

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/css"
)

//Data stores html data with a seed.
type Data struct {
	ID  *string
	Tag string

	Classes []string

	InnerHTML string

	Style map[string]string

	Attributes map[string]string
}

//SetID returns an option that sets the HTML id associated with the seed.
func SetID(id string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Load(&data)

		switch mode, q := client.Seed(c); mode {
		case client.AddTo:
			fmt.Fprintf(q, `%v.id = %v;`, client.Element(c), strconv.Quote(id))
		case client.Undo:
			if data.ID != nil {
				fmt.Fprintf(q, `%v.id = %v;`, client.Element(c), strconv.Quote(*data.ID))
			} else {
				fmt.Fprintf(q, `%v.id = %v;`, client.Element(c), c.ID())
			}
		default:
			data.ID = &id
		}

		c.Save(data)
	})
}

//AddClass adds a class to the html element.
func AddClass(class string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Load(&data)

		switch mode, q := client.Seed(c); mode {
		case client.AddTo:
			fmt.Fprintf(q, `%v.classList.With(%v);`, client.Element(c), strconv.Quote(class))
		case client.Undo:
			fmt.Fprintf(q, `%v.classList.remove(%v);`, client.Element(c), strconv.Quote(class))
		default:

			for _, existing := range data.Classes {
				if class == existing {
					c.Save(data)
					return
				}
			}

			data.Classes = append(data.Classes, class)
		}

		c.Save(data)
	})
}

//SetTag returns an option that sets the HTML tag associated with the seed.
func SetTag(tag string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Load(&data)

		switch mode, q := client.Seed(c); mode {
		case client.AddTo:
			fmt.Fprintf(q, `%v = document.createElement("%v"); %v.id = "temp%v";`, client.Element(c), tag, client.Element(c), c.ID())
		case client.Undo:
			fmt.Fprintf(q, `%v = document.createElement("%v"); %v.id = "temp%v";`, client.Element(c), data.Tag, client.Element(c), c.ID())
		default:
			data.Tag = tag
		}

		c.Save(data)
	})
}

//Set returns an option that sets the HTML associated with the seed.
func Set(html string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Load(&data)

		switch mode, q := client.Seed(c); mode {
		case client.AddTo:
			fmt.Fprintf(q, `%v.innerHTML = %v;`, client.Element(c), strconv.Quote(html))
		case client.Undo:
			fmt.Fprintf(q, `%v.innerHTML = %v;`, client.Element(c), strconv.Quote(data.InnerHTML))
		default:
			data.InnerHTML = html
		}

		c.Save(data)
	})
}

//SetAttribute returns an option that sets an HTML attribute of this seed.
func SetAttribute(name string, constant string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Load(&data)

		switch mode, q := client.Seed(c); mode {
		case client.AddTo:
			fmt.Fprintf(q, `%v.setAttribute(%v, %v);`, client.Element(c), strconv.Quote(name), strconv.Quote(constant))
		case client.Undo:
			if attr, ok := data.Attributes[name]; ok {
				fmt.Fprintf(q, `%v.setAttribute(%v, %v);`, client.Element(c), strconv.Quote(name), strconv.Quote(attr))
			} else {
				fmt.Fprintf(q, `%v.removeAttribute(%v);`, client.Element(c), strconv.Quote(name))
			}
		default:
			if data.Attributes == nil {
				data.Attributes = make(map[string]string)
			}
			data.Attributes[name] = constant
		}

		c.Save(data)
	})
}

//SetAttributeTo returns an option that sets an HTML attribute of this seed.
func SetAttributeTo(name string, variable client.String) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		if variable != nil {
			clientside.Hook(variable, c)
			client.OnRender(
				Element(c).Run("setAttribute", client.NewString(name), variable.GetString()),
			).AddTo(c)
		}
	})
}

//SetStyle returns an option that sets the inline HTML style of the seed.
func SetStyle(property, value string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Load(&data)

		switch mode, _ := client.Seed(c); mode {
		case client.AddTo, client.Undo:
			css.Set(property, value).AddTo(c)
		default:
			if data.Style == nil {
				data.Style = make(map[string]string)
			}
			data.Style[property] = value
		}

		c.Save(data)
	})
}

//SetInnerText returns an option that sets the HTML innerText associated with the seed.
func SetInnerText(constant string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {

		var data Data
		c.Load(&data)

		switch mode, q := client.Seed(c); mode {
		case client.AddTo:
			fmt.Fprintf(q, `%v.innerText = %v;`, client.Element(c), strconv.Quote(constant))
		case client.Undo:
			fmt.Fprintf(q, `%v.innerHTML = %v;`, client.Element(c), strconv.Quote(data.InnerHTML))
		default:
			data.InnerHTML = html.EscapeString(constant)
			data.InnerHTML = strings.Replace(data.InnerHTML, "\n", "<br>", -1)
			data.InnerHTML = strings.Replace(data.InnerHTML, "  ", "&nbsp;&nbsp;", -1)
			data.InnerHTML = strings.Replace(data.InnerHTML, "\t", "&emsp;", -1)
		}

		c.Save(data)
	})
}

//SetInnerTextTo returns an option that sets the HTML innerText associated with the seed.
func SetInnerTextTo(variable client.String) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		if variable != nil {
			clientside.Hook(variable, c)
			client.OnRender(
				Element(c).Set("innerText", variable.GetString()),
			).AddTo(c)
		}
	})
}
