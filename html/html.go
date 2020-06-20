package html

import (
	"html"
	"strconv"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/script"
)

//Data stores html data with a seed.
type Data struct {
	seed.Data

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
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.id = %v;`, q.Element(), strconv.Quote(id))
		case script.Undo:
			if data.ID != nil {
				q.Javascript(`%v.id = %v;`, q.Element(), strconv.Quote(*data.ID))
			} else {
				q.Javascript(`%v.id = %v;`, q.Element(), c.ID())
			}
		default:
			data.ID = &id
		}

		c.Write(data)
	})
}

//AddClass adds a class to the html element.
func AddClass(class string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.classList.With(%v);`, q.Element(), strconv.Quote(class))
		case script.Undo:
			q.Javascript(`%v.classList.remove(%v);`, q.Element(), strconv.Quote(class))
		default:

			for _, existing := range data.Classes {
				if class == existing {
					c.Write(data)
					return
				}
			}

			data.Classes = append(data.Classes, class)
		}

		c.Write(data)
	})
}

//SetTag returns an option that sets the HTML tag associated with the seed.
func SetTag(tag string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v = document.createElement("%v"); %v.id = "temp%v";`, q.Element(), tag, q.Element(), c.ID())
		case script.Undo:
			q.Javascript(`%v = document.createElement("%v"); %v.id = "temp%v";`, q.Element(), data.Tag, q.Element(), c.ID())
		default:
			data.Tag = tag
		}

		c.Write(data)
	})
}

//Set returns an option that sets the HTML associated with the seed.
func Set(html string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.innerHTML = %v;`, q.Element(), strconv.Quote(html))
		case script.Undo:
			q.Javascript(`%v.innerHTML = %v;`, q.Element(), strconv.Quote(data.InnerHTML))
		default:
			data.InnerHTML = html
		}

		c.Write(data)
	})
}

//SetAttribute returns an option that sets an HTML attribute of this seed.
func SetAttribute(name, value string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.setAttribute(%v, %v);`, q.Element(), strconv.Quote(name), strconv.Quote(value))
		case script.Undo:
			if attr, ok := data.Attributes[name]; ok {
				q.Javascript(`%v.setAttribute(%v, %v);`, q.Element(), strconv.Quote(name), strconv.Quote(attr))
			} else {
				q.Javascript(`%v.removeAttribute(%v);`, q.Element(), strconv.Quote(name))
			}
		default:
			if data.Attributes == nil {
				data.Attributes = make(map[string]string)
			}
			data.Attributes[name] = value
		}

		c.Write(data)
	})
}

//SetStyle returns an option that sets the inline HTML style of the seed.
func SetStyle(property, value string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data Data
		c.Read(&data)

		switch c.(type) {
		case script.Seed, script.Undo:
			css.Set(property, value).AddTo(c)
		default:
			if data.Style == nil {
				data.Style = make(map[string]string)
			}
			data.Style[property] = value
		}

		c.Write(data)
	})
}

//SetInnerText returns an option that sets the HTML innerText associated with the seed.
func SetInnerText(text string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {

		var data Data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.innerText = %v;`, q.Element(), strconv.Quote(text))
		case script.Undo:
			q.Javascript(`%v.innerHTML = %v;`, q.Element(), strconv.Quote(data.InnerHTML))
		default:
			data.InnerHTML = html.EscapeString(text)
			data.InnerHTML = strings.Replace(data.InnerHTML, "\n", "<br>", -1)
			data.InnerHTML = strings.Replace(data.InnerHTML, "  ", "&nbsp;&nbsp;", -1)
			data.InnerHTML = strings.Replace(data.InnerHTML, "\t", "&emsp;", -1)
		}

		c.Write(data)
	})
}
