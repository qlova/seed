package feed

import (
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientop"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
	"qlova.org/seed/s/html/div"
	"qlova.org/seed/s/html/template"
)

//Field can be used to select feed data.
type Field interface {
	FieldName() string
}

type field string

func (f field) FieldName() string {
	return string(f)
}

//Get converts a string into a field.
func Get(name string) Field {
	return field(name)
}

//Feed has food ready to populate.
type Feed struct {
	feed, template seed.Seed

	food Food

	Data Item

	//boolean is true if the feed has items.
	boolean *clientside.Bool

	Empty clientop.Bool
}

//GetBool implements js.AnyBool
func (f *Feed) GetBool() js.Bool {
	return f.boolean.GetBool()
}

//GetValue implements js.AnyValue
func (f *Feed) GetValue() js.Value {
	return f.GetBool().Value
}

//Components implements client.Compound
func (f *Feed) Components() []client.Value {
	return []client.Value{f.boolean}
}

func (f *Feed) String(field Field) client.String {
	return js.String{Value: f.Data.Get(field.FieldName())}
}

func (f *Feed) Int(field Field) client.Int {
	return js.Number{Value: f.Data.Get(field.FieldName())}
}

//Refresh refreshes the feed.
func (f *Feed) Refresh() client.Script {
	return html.Element(f.template).Run("onrefresh")
}

//With returns a new Feed on the given Food, the options provided will be applied to the feed itself.
func With(food Food, options ...seed.Option) *Feed {

	var template = template.New()
	var feed = div.New(template,
		css.Set("display", "flex"),
		css.Set("flex-direction", "column"),
		seed.Options(options),
	)

	var f = &Feed{
		feed: feed,
		food: food,

		template: template,

		boolean: new(clientside.Bool),

		Data: Item{
			array: js.Array{js.NewValue("q.feed")},
			Value: js.NewValue("q.data"),
			Index: js.Number{js.NewValue("q.i")},
		},
	}

	f.Empty = clientop.Not(f.boolean)

	return f
}

//New returns a new instantiated feed with the given options.
func (f *Feed) New(options ...seed.Option) seed.Seed {
	var template = f.template
	var feed = f.feed

	template.With(css.SetSelector("#"+html.ID(feed)), seed.Options(options))

	convertToClasses(template)

	var scripts js.Script = func(js.Ctx) {}
	for _, child := range template.Children() {
		scripts = scripts.Append(client.Adopt(child).GetScript())
	}

	var scriptsString strings.Builder
	js.NewCtx(&scriptsString)(scripts)

	mem, adr := f.boolean.Variable()

	feed.With(
		client.OnLoad(js.Func("s.feed.orf").Run(js.NewValue("q"), js.NewString(client.ID(feed)), js.NewFunction(func(q js.Ctx) {
			q.Return(food2Data(f.food, q))
		}), js.NewFunction(func(q js.Ctx) {
			q("return async function(q) {")
			q(scripts)
			q("};")
		}), js.NewString(string(mem)), js.NewString(string(adr)))),
	)

	return feed
}
