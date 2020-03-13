package page

import (
	"fmt"
	"reflect"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style"
)

type harvester struct {
	Parent seed.Any

	Map map[reflect.Type]seed.Seed
}

func newHarvester(parent seed.Any) harvester {
	return harvester{
		Parent: parent,
		Map:    make(map[reflect.Type]seed.Seed),
	}
}

func (h harvester) harvest(page Page) seed.Seed {
	key := reflect.TypeOf(page)
	if h.Map[key] == 0 {

		//Harvest the page
		var template = seed.New(
			html.SetTag("template"),
		)
		template.Use()
		template.AddTo(h.Parent)

		var element = seed.New(
			html.SetTag("div"),
			css.SetDisplay(css.Flex),
			css.SetFlexDirection(css.Column),

			style.SetSize(100, 100),
		)
		element.Use()
		element.AddTo(template)

		h.Map[key] = element.Root()

		page.Page(Scope{element, h})
	}

	return h.Map[key]
}

//Harvest returns an option that adds all pages to the acting seed.
//This should normally only be called by app-level runtime packages such as seed/app.
func Harvest(p Page) seed.Option {
	return seed.Do(func(c seed.Seed) {
		if p == nil {
			return
		}

		id := newHarvester(c).harvest(p)

		c.Add(script.OnReady(func(q script.Ctx) {
			fmt.Fprintf(q, `seed.goto.ready("%v");`, html.ID(id))
		}))
	})
}
