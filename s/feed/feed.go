package feed

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"

	"github.com/qlova/seed/s/html/div"
	"github.com/qlova/seed/s/html/template"
)

type Data struct {
	string
}

func (d Data) String() state.String {
	return state.String{Value: state.Raw(d.string, state.Local())}
}

func (d Data) Get(name string) Data {
	return Data{fmt.Sprintf(`%v[%q]`, d.string, name)}
}

type Seed struct {
	seed.Seed
	Data Data
}

func (c Seed) Refresh() script.Script {
	return func(q script.Ctx) {
		fmt.Fprintf(q, `await %v.refresh();`, script.Scope(c, q).Element())
	}
}

func Refresh(c seed.Seed) script.Script {
	return func(q script.Ctx) {
		fmt.Fprintf(q, `await %v.refresh();`, script.Scope(c, q).Element())
	}
}

//Do runs f.
func Do(f func(Seed)) seed.Option {
	return seed.Do(func(s seed.Seed) {
		f(Seed{s, Data{"data"}})
	})
}

func convertToClasses(c seed.Seed) {
	for _, child := range c.Children() {
		child.Add(
			css.SetSelector(`.`+html.ID(child)),
			html.SetID(""),
			html.AddClass(html.ID(child)),
		)

		convertToClasses(child)
	}
}

//Food is fed to a feed to populate it with items.
type Food interface{}

type rpc struct {
	f    interface{}
	args []script.AnyValue
}

func Go(f interface{}, args ...script.AnyValue) Food {
	return rpc{f, args}
}

func food2Data(food Food, q script.Ctx) script.Value {
	switch reflect.TypeOf(food).Kind() {
	case reflect.Func:
		return script.RPC(food)(q)
	default:
		switch f := food.(type) {
		case rpc:
			return script.RPC(f.f, f.args...)(q)
		}
		panic("unsupported data type")
	}
}

//New returns a repeater capable of repeating itself based on the given Go data.
func New(food Food, options ...seed.Option) Seed {
	var template = template.New()
	var feed = Seed{div.New(template,
		css.Set("display", "flex"),
		css.Set("flex-direction", "column"),
	), Data{}}

	template.Add(css.SetSelector("#"+html.ID(feed.Seed)), seed.Options(options))

	convertToClasses(template)
	var scripts script.Script = func(js.Ctx) {}
	for _, child := range template.Children() {
		scripts = scripts.Append(script.Adopt(child))
	}

	feed.Add(
		script.OnReady(func(q script.Ctx) {
			fmt.Fprintf(q, `%v.refresh = async function() {
				try {
			if (%[1]v.refreshing) return;
			%[1]v.refreshing = true; 
			let cache = seed.get.cache; 
			seed.get.cache = null;`, js.NewValue(script.Scope(template, q).Element()))

			fmt.Fprintf(q, `while (%[1]v.childNodes.length > 1) %[1]v.removeChild(%[1]v.lastChild);`, script.Scope(feed, q).Element())

			var data = food2Data(food, q)

			var scriptsString strings.Builder
			scriptsString.WriteString(`async function(data) {`)
			js.NewCtx(&scriptsString)(scripts)
			scriptsString.WriteString(`}`)

			q.Run(`await seeds.feed.refresh`, js.NewValue(script.Scope(template, q).Element()), data, js.NewValue(scriptsString.String()))
			fmt.Fprintf(q, `seed.get.cache = cache; %v.refreshing = false; }
				catch(e) {
					%[1]v.refreshing = false;
					throw e;
				}
			};`, script.Scope(template, q).Element())
		}))

	return feed
}
