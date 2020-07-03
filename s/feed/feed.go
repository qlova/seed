package feed

import (
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
	"qlova.org/seed/s/html/div"
	"qlova.org/seed/s/html/template"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
)

//Field can be used to select feed data.
type Field interface {
	FieldName() string
}

//Feed has food ready to populate.
type Feed struct {
	feed, template seed.Seed

	food Food

	Data Item
}

func (f Feed) String(field Field) client.String {
	return js.String{Value: f.Data.Get(field.FieldName())}
}

func (f Feed) Int(field Field) client.Int {
	return js.Number{Value: f.Data.Get(field.FieldName())}
}

//Refresh refreshes the feed.
func (f Feed) Refresh() client.Script {
	return script.Element(f.template).Run("onrefresh")
}

//With returns a new Feed on the given Food, the options provided will be applied to the feed itself.
func With(food Food, options ...seed.Option) Feed {

	var template = template.New()
	var feed = div.New(template,
		css.Set("display", "flex"),
		css.Set("flex-direction", "column"),
		seed.Options(options),
	)

	return Feed{
		feed: feed,
		food: food,

		template: template,

		Data: Item{
			array: js.Array{js.NewValue("q.feed")},
			Value: js.NewValue("q.data"),
			Index: js.Number{js.NewValue("q.i")},
		},
	}
}

//New returns a new instantiated feed with the given options.
func (f Feed) New(options ...seed.Option) seed.Seed {
	var template = f.template
	var feed = f.feed

	template.With(css.SetSelector("#"+html.ID(feed)), seed.Options(options))

	convertToClasses(template)

	var scripts script.Script = func(js.Ctx) {}
	for _, child := range template.Children() {
		scripts = scripts.Append(script.Adopt(child))
	}

	var rerender script.Script = state.AdoptRefreshOfChildren(template)

	var scriptsString strings.Builder
	js.NewCtx(&scriptsString)(scripts)
	js.NewCtx(&scriptsString)(rerender)

	feed.With(
		script.OnReady(js.Func("s.feed.orf").Run(js.NewValue("q"), js.NewString(script.ID(feed)), js.NewFunction(func(q script.Ctx) {
			q.Return(food2Data(f.food, q))
		}), js.NewFunction(func(q script.Ctx) {
			q("return async function(q) {")
			q(scripts)
			q(rerender)
			q("};")
		}))),
	)

	return feed
}
