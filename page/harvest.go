package page

import (
	"fmt"
	"reflect"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
	"qlova.org/seed/s/column"
	"qlova.org/seed/s/expander"
	"qlova.org/seed/set"
	"qlova.org/seed/units/percentage/of"
)

type harvester struct {
	Parent seed.Seed

	Map map[reflect.Type]seed.Seed
}

func newHarvester(parent seed.Seed) harvester {
	return harvester{
		Parent: parent,
		Map:    make(map[reflect.Type]seed.Seed),
	}
}

func Set(page Page) seed.Option {
	return client.OnLoad(js.Script(func(q js.Ctx) {
		fmt.Fprintf(q, `seed.StartingPage = "%v";`, ID(page))
	}))
}

func AddPages(pages ...Page) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var container = column.New(
			set.Width(100%of.Parent),
			expander.Set(),
			set.MinHeight(nil),
		)
		for _, page := range pages {
			add(page).AddTo(container)
		}
		container.AddTo(c)
	})
}

func add(page Page) seed.Option {
	if page == nil {
		panic("page is nil")
	}
	return seed.NewOption(func(c seed.Seed) {
		var template = seed.New(
			html.SetTag("template"),
		)
		template.Use()
		template.AddTo(c)

		page, _, _ = parseArgs(page)

		var element = page.Page(RouterOf(template))
		element.With(
			html.AddClass("page"),
			html.AddClass(ID(page)[1:]),
			css.SetSelector(ID(page)),
			client.SetID(ID(page)),
			html.SetID(html.ID(element)),
		)
		element.Use()

		element.AddTo(template)
	})
}

func (h harvester) harvest(c seed.Seed) {
	var data data
	c.Load(&data)

	for _, p := range data.pages {
		var key = reflect.TypeOf(p)

		if _, ok := h.Map[key]; !ok {
			add(p).AddTo(h.Parent)

			h.Map[key] = c
		}
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}
}

//Harvest returns an option that adds all pages to the acting seed.
//This should normally only be called by app-level runtime packages such as seed/app.
func Harvest(starting Page) seed.Option {
	if starting == nil {
		return seed.NewOption(func(c seed.Seed) {})
	}
	return seed.NewOption(func(c seed.Seed) {
		if starting == nil {
			return
		}

		var data data
		c.Load(&data)
		data.pages = append(data.pages, starting)
		c.Save(data)

		var container = column.New(
			set.Width(100%of.Parent),
			expander.Set(),
		)
		newHarvester(container).harvest(c)
		container.AddTo(c)

		c.With(client.OnLoad(js.Script(func(q js.Ctx) {
			fmt.Fprintf(q, `seed.goto.ready("%v");`, ID(starting))
		})))
	})
}
