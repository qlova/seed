package html

import (
	"html"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/script"
)

type data struct {
	seed.Data

	id  *string
	tag string

	classes []string

	innerHTML string

	style map[string]string

	attributes map[string]string
}

var seeds = make(map[seed.Seed]data)

//SetID returns an option that sets the HTML id associated with the seed.
func SetID(id string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.id = %v;`, q.Element(), strconv.Quote(id))
		case script.Undo:
			if data.id != nil {
				q.Javascript(`%v.id = %v;`, q.Element(), strconv.Quote(*data.id))
			} else {
				q.Javascript(`%v.id = %v;`, q.Element(), c.ID())
			}
		default:
			data.id = &id
		}

		c.Write(data)
	})
}

//AddClass adds a class to the html element.
func AddClass(class string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.classList.add(%v);`, q.Element(), strconv.Quote(class))
		case script.Undo:
			q.Javascript(`%v.classList.remove(%v);`, q.Element(), strconv.Quote(class))
		default:
			data.classes = append(data.classes, class)
		}

		c.Write(data)
	})
}

//SetTag returns an option that sets the HTML tag associated with the seed.
func SetTag(tag string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v = document.createElement("%v"); %v.id = "temp%v";`, q.Element(), tag, q.Element(), c.ID())
		case script.Undo:
			q.Javascript(`%v = document.createElement("%v"); %v.id = "temp%v";`, q.Element(), data.tag, q.Element(), c.ID())
		default:
			data.tag = tag
		}

		c.Write(data)
	})
}

//Set returns an option that sets the HTML associated with the seed.
func Set(html string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.innerHTML = %v;`, q.Element(), strconv.Quote(html))
		case script.Undo:
			q.Javascript(`%v.innerHTML = %v;`, q.Element(), strconv.Quote(data.innerHTML))
		default:
			data.innerHTML = html
		}

		c.Write(data)
	})
}

//SetAttribute returns an option that sets an HTML attribute of this seed.
func SetAttribute(name, value string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.setAttribute(%v, %v);`, q.Element(), strconv.Quote(name), strconv.Quote(value))
		case script.Undo:
			if attr, ok := data.attributes[name]; ok {
				q.Javascript(`%v.setAttribute(%v, %v);`, q.Element(), strconv.Quote(name), strconv.Quote(attr))
			} else {
				q.Javascript(`%v.removeAttribute(%v);`, q.Element(), strconv.Quote(name))
			}
		default:
			if data.attributes == nil {
				data.attributes = make(map[string]string)
			}
			data.attributes[name] = value
		}

		c.Write(data)
	})
}

//SetStyle returns an option that sets the inline HTML style of the seed.
func SetStyle(property, value string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Read(&data)

		switch c.(type) {
		case script.Seed, script.Undo:
			css.Set(property, value).AddTo(c)
		default:
			if data.style == nil {
				data.style = make(map[string]string)
			}
			data.style[property] = value
		}

		c.Write(data)
	})
}

//SetInnerText returns an option that sets the HTML innerText associated with the seed.
func SetInnerText(text string) seed.Option {
	return seed.NewOption(func(c seed.Seed) {

		var data data
		c.Read(&data)

		switch q := c.(type) {
		case script.Seed:
			q.Javascript(`%v.innerText = %v;`, q.Element(), strconv.Quote(text))
		case script.Undo:
			q.Javascript(`%v.innerHTML = %v;`, q.Element(), strconv.Quote(data.innerHTML))
		default:
			data.innerHTML = html.EscapeString(text)
		}

		c.Write(data)
	})
}
