package feed

import (
	"qlova.org/seed/client"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

//Item is an individual element of a feed.
type Item struct {
	js.Value

	Index client.Int
	array js.Array
}

func (item Item) Previous() Item {
	item.Value = js.NewValue("(%v || {})", item.array.Index(item.Index.GetNumber().Minus(js.NewNumber(1))))
	return item
}

//Filter filters the food on the client with a given filter function.
func Filter(food Food, fn func(Item) client.Bool) Food {
	return func(q script.Ctx) js.Value {
		return js.NewValue(`%v.filter(%v)`, food2Data(food, q), js.NewNormalFunction(func(q script.Ctx) {
			q.Return(fn(Item{
				Value: js.NewValue("value"),
				Index: js.Number{js.NewValue("index")},
				array: js.Array{js.NewValue("array")},
			}))
		}, "value", "index", "array"))
	}
}
