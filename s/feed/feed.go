package feed

import (
	"fmt"
	"reflect"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/tween"

	"github.com/qlova/seed/s/html/div"
	"github.com/qlova/seed/s/html/template"
)

type Data struct {
	string
}

func (d Data) String() state.String {
	return state.String{Value: state.Raw(`data`, state.Local())}
}

type Seed struct {
	seed.Seed
	Data Data
}

func (c Seed) Refresh() script.Script {
	return func(q script.Ctx) {
		fmt.Fprintf(q, `%v.refresh();`, q.Scope(c).Element())
	}
}

//Do runs f.
func Do(f func(Seed)) seed.Option {
	return seed.Do(func(s seed.Seed) {
		f(Seed{s, Data{}})
	})
}

func convertToClasses(c seed.Seed) {
	for _, child := range c.Children() {
		child.Add(css.SetSelector(`.`+html.ID(child)), html.SetID(""), html.AddClass(html.ID(child)))
	}
}

//Food is fed to a feed to populate it with items.
type Food interface{}

//New returns a repeater capable of repeating itself based on the given Go data.
func New(food Food, options ...seed.Option) Seed {
	var template = template.New()
	var feed = Seed{div.New(template,
		css.Set("display", "flex"),
		css.Set("flex-direction", "column"),
	), Data{}}

	feed.Add(tween.This())

	template.Add(css.SetSelector("#" + html.ID(feed.Seed)).And(options...))

	convertToClasses(template)
	var scripts script.Script
	for _, child := range template.Children() {
		scripts = scripts.Then(script.Adopt(child))
	}

	switch reflect.TypeOf(food).Kind() {
	case reflect.Func:

	default:
		panic("unsupported data type")
	}

	feed.Add(script.OnReady(func(q script.Ctx) {
		fmt.Fprintf(q, `%v.refresh = async function() {`, q.Scope(template).Element())

		fmt.Fprintf(q, `while (%[1]v.childNodes.length > 1) %[1]v.removeChild(%[1]v.lastChild);`, q.Scope(feed).Element())

		var data = q.Go(food).Wait().Native
		fmt.Fprintf(q, `if (!Array.isArray(%[1]v)) %[1]v = [%[1]v];`, q.Raw(data))
		fmt.Fprintf(q, `for (let value of %v) {`, q.Raw(data))
		{
			fmt.Fprintf(q, `let data = value; let clone = %v.content.cloneNode(true);`, q.Scope(template).Element())

			fmt.Fprintf(q, `let cache = seed.get.cache; let old = seed.get; let parent = %v; seed.get = function(id) {
				let get = old(id);
				if (!get) return clone.querySelector("."+id);
				return get;
			};seed.get.cache = cache; `, q.Scope(feed).Element())
			scripts(q)
			fmt.Fprintf(q, `seed.get = old;`)
			fmt.Fprintf(q, `clone = %v.appendChild(clone)`, q.Scope(feed).Element())
		}
		fmt.Fprintf(q, `}};`)
	}))

	return feed
}
