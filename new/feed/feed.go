package feed

import (
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/client/if/not"
	"qlova.org/seed/web/css"
	"qlova.org/seed/web/html"
	"qlova.org/seed/web/js"
	"qlova.org/seed/new/html/div"
	"qlova.org/seed/new/html/template"
)

type data struct {
	

	templates []seed.Seed
}

func Templates(root seed.Seed) []seed.Seed {
	var data data
	root.Load(&data)

	var result = data.templates

	for _, child := range root.Children() {
		if slice := Templates(child); slice != nil {
			result = append(result, slice...)
		}
	}

	return result
}

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

	Empty client.Bool
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
	return html.Element(f.feed).Run("onrefresh")
}

//With returns a new Feed on the given Food, the options provided will be applied to the feed itself.
func With(food Food, options ...seed.Option) *Feed {

	var template = template.New()
	var feed = div.New(
		css.Set("display", "flex"),
		css.Set("flex-direction", "column"),

		seed.Mutate(func(data *data) {
			data.templates = append(data.templates, template)
		}),

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

	f.Empty = not.True(f.boolean)

	return f
}

//New returns a new instantiated feed with the given options.
func (f *Feed) New(options ...seed.Option) seed.Seed {
	var template = f.template
	var feed = f.feed

	template.With(css.SetSelector("#"+html.ID(feed)), html.AddClass("sortable-ignore"), seed.Options(options))

	convertToClasses(template)

	var scripts js.Script = func(js.Ctx) {}
	for _, child := range template.Children() {
		scripts = scripts.Append(client.Adopt(child).GetScript())
	}

	var scriptsString strings.Builder
	js.NewCtx(&scriptsString)(scripts)

	mem, adr := f.boolean.Variable()

	feed.With(
		client.OnLoad(js.Func("s.feed.orf").Run(js.NewValue("q"), js.NewString(client.ID(feed)), js.NewString(client.ID(template)), js.NewFunction(func(q js.Ctx) {
			q.Return(food2Data(f.food, q))
		}), js.NewFunction(func(q js.Ctx) {
			q("return async function(q) {")
			q(scripts)
			q("};")
		}), js.NewString(string(mem)), js.NewString(string(adr)))),
	)

	return feed
}
