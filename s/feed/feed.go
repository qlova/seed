package feed

import (
	"fmt"
	"reflect"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/style"

	"qlova.org/seed/s/html/div"
	"qlova.org/seed/s/html/template"
)

type Data struct {
	string
}

func (d Data) String() state.String {
	return state.String{Value: state.Raw("("+d.string+` || "")`, state.Local())}
}

func (d Data) Number() js.Number {
	return js.Number{js.NewValue("(" + d.string + ` || 0)`)}
}

func (d Data) Get(name string) Data {
	return Data{fmt.Sprintf(`%v[%q]`, d.string, name)}
}

type Seed struct {
	seed.Seed
	Data Data
}

func (c Seed) Refresh() script.Script {
	return script.Element(c).Run("onrefresh")
}

func Refresh(c seed.Seed) script.Script {
	return script.Element(c).Run("onrefresh")
}

//Do runs f.
func Do(f func(Seed)) seed.Option {
	return seed.NewOption(func(s seed.Seed) {
		f(Seed{s, Data{`q.data`}})
	})
}

func convertToClasses(c seed.Seed) {
	if _, ok := c.(Seed); ok {
		return
	}

	for _, child := range c.Children() {
		child.With(
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
		case script.Value:
			return f
		case script.AnyValue:
			return f.GetValue()
		}
		panic("unsupported feed.Food: " + reflect.TypeOf(food).String())
	}
}

//New returns a repeater capable of repeating itself based on the given Go data.
func New(food Food, options ...seed.Option) Seed {
	var styles seed.Options
	var others seed.Options

	for _, o := range options {
		if _, ok := o.(style.Style); ok {
			styles = append(styles, o)
		} else {
			others = append(others, o)
		}
	}

	var template = template.New()
	var feed = Seed{div.New(template,
		css.Set("display", "flex"),
		css.Set("flex-direction", "column"),
		styles,
	), Data{}}

	template.With(css.SetSelector("#"+html.ID(feed.Seed)), others)

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
			q.Return(food2Data(food, q))
		}), js.NewFunction(func(q script.Ctx) {
			q("return async function(q) {")
			q(scripts)
			q(rerender)
			q("};")
		}))),
	)

	return feed
}
