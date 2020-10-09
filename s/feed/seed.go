package feed

import (
	"reflect"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
)

func convertToClasses(c seed.Seed) {
	for _, child := range c.Children() {

		if sc, ok := child.(client.Seed); ok {
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
	args []client.Value
}

func Go(f interface{}, args ...client.Value) Food {
	return rpc{f, args}
}

func food2Data(food Food, q js.Ctx) client.Value {
	if food == nil {
		return js.Null()
	}
	switch reflect.TypeOf(food).Kind() {
	case reflect.Func:
		switch f := food.(type) {
		case func(q js.Ctx) js.Value:
			return f(q)
		}
		return client.Call(food)
	default:
		switch f := food.(type) {
		case rpc:
			return client.Call(f.f, f.args...)
		case client.Value:
			return f
		}
		panic("unsupported feed.Food: " + reflect.TypeOf(food).String())
	}
}
