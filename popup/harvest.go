package popup

import (
	"reflect"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/script"
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

func (h harvester) harvest(c seed.Seed) {
	var data data
	c.Read(&data)

	for _, p := range data.popups {
		var key = reflect.TypeOf(p)
		if _, ok := h.Map[key]; !ok {

			var template = seed.New(
				html.SetTag("template"),
			)
			template.Use()
			template.AddTo(h.Parent)

			var element = p.Popup(ManagerOf(template))
			element.With(
				html.SetID(ID(p)),
				script.SetID(ID(p)),
				css.SetSelector("#"+ID(p)),
			)
			element.Use()
			element.AddTo(template)

			h.Map[key] = element

			h.harvest(element)
		}
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}
}

//Harvest returns an option that adds all pages to the acting seed.
//This should normally only be called by app-level runtime packages such as seed/app.
func Harvest() seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		newHarvester(c).harvest(c)
	})
}
