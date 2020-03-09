package page

import (
	"fmt"
	"reflect"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/script"
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
		)
		element.Use()
		element.AddTo(template)

		h.Map[key] = element.Root()

		page.Page(Seed{element, h})
	}

	return h.Map[key]
}

func Set(current Page) seed.Option {
	return seed.NewOption(func(any seed.Any) {
		var h = newHarvester(any)

		var id = h.harvest(current)

		any.Add(script.OnReady(func(q script.Ctx) {
			fmt.Fprintf(q, `seed.goto.ready("%v");`, html.ID(id))
		}))

	}, func(ctx seed.Ctx) {

	}, func(ctx seed.Ctx) {

	})
}

type Page interface {
	Page(Seed)
}

type Seed struct {
	seed.Seed
	harvester
}

func (seed Seed) Goto(page Page) script.Script {
	return func(q script.Ctx) {
		seed.Ctx(q).Goto(page)
	}
}

func OnEnter(f script.Script) seed.Option {
	return script.On("pageenter", f)
}

func OnExit(f script.Script) seed.Option {
	return script.On("pageexit", f)
}

type Ctx struct {
	Seed
	seed.Ctx
}

func (seed Seed) Ctx(q script.AnyCtx) Ctx {
	return Ctx{seed, seed.Seed.Ctx(q)}
}

func (s Ctx) Goto(page Page) {
	fmt.Fprintf(s.Ctx.Ctx, `seed.goto("%v");`, html.ID(s.Seed.harvester.harvest(page)))
}
