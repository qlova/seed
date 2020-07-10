package feed

import (
	"reflect"

	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

func convertToClasses(c seed.Seed) {
	for _, child := range c.Children() {

		if sc, ok := child.(script.Seed); ok {
			child = sc.Seed
		}

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
	if food == nil {
		return js.Null()
	}
	switch reflect.TypeOf(food).Kind() {
	case reflect.Func:
		switch f := food.(type) {
		case func(q script.Ctx) js.Value:
			return f(q)
		}
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
