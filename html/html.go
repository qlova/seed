package html

import (
	"fmt"
	"html"
	"strconv"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
)

type data struct {
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
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		data.id = &id
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v.id = %v;`, s.Element(), strconv.Quote(id))
	}, func(s seed.Ctx) {
		data := seeds[s.Root()]
		if data.id != nil {
			fmt.Fprintf(s.Ctx, `%v.id = %v;`, s.Element(), strconv.Quote(*data.id))
		} else {
			fmt.Fprintf(s.Ctx, `%v.id = %v;`, s.Element(), s.Root())
		}
	})
}

//AddClass adds a class to the html element.
func AddClass(class string) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		data.classes = append(data.classes, class)
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v.classList.add(%v);`, s.Element(), strconv.Quote(class))
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v.classList.remove(%v);`, s.Element(), strconv.Quote(class))
	})
}

//SetTag returns an option that sets the HTML tag associated with the seed.
func SetTag(tag string) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		data.tag = tag
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v = document.createElement("%v"); %v.id = "temp%v";`, s.Element(), tag, s.Element(), s.Root())
	}, func(s seed.Ctx) {
		data := seeds[s.Root()]
		fmt.Fprintf(s.Ctx, `%v = document.createElement("%v"); %v.id = "temp%v";`, s.Element(), data.tag, s.Element(), s.Root())
	})
}

//Set returns an option that sets the HTML associated with the seed.
func Set(html string) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		data.innerHTML = html
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v.innerHTML = %v;`, s.Element(), strconv.Quote(html))
	}, func(s seed.Ctx) {
		data := seeds[s.Root()]
		fmt.Fprintf(s.Ctx, `%v.innerHTML = %v;`, s.Element(), strconv.Quote(data.innerHTML))
	})
}

//SetAttribute returns an option that sets an HTML attribute of this seed.
func SetAttribute(name, value string) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		if data.attributes == nil {
			data.attributes = make(map[string]string)
		}
		data.attributes[name] = value
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v.setAttribute(%v, %v)`, s.Element(), strconv.Quote(name), strconv.Quote(value))
	}, func(s seed.Ctx) {
		data := seeds[s.Root()]
		fmt.Fprintf(s.Ctx, `%v.setAttribute(%v, %v)`, s.Element(), strconv.Quote(name), data.attributes[name])
	})
}

//SetStyle returns an option that sets the inline HTML style of the seed.
func SetStyle(property, value string) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		if data.style == nil {
			data.style = make(map[string]string)
		}
		data.style[property] = value
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		css.Set(property, value).Apply(s)
	}, func(s seed.Ctx) {
		css.Set(property, value).Reset(s)
	})
}

//SetInnerText returns an option that sets the HTML innerText associated with the seed.
func SetInnerText(text string) seed.Option {
	return seed.NewOption(func(s seed.Any) {
		data := seeds[s.Root()]
		data.innerHTML = html.EscapeString(text)
		seeds[s.Root()] = data
	}, func(s seed.Ctx) {
		fmt.Fprintf(s.Ctx, `%v.innerText = %v;`, s.Element(), strconv.Quote(text))
	}, func(s seed.Ctx) {
		data := seeds[s.Root()]
		fmt.Fprintf(s.Ctx, `%v.innerHTML = %v;`, s.Element(), strconv.Quote(data.innerHTML))
	})
}
