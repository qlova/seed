//go:generate go run -tags generate generate.go

package sortable

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/change"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/client/if/not"
	"qlova.org/seed/web/html"
	"qlova.org/seed/web/html/attr"
	"qlova.org/seed/web/js"
)

type data struct {
	Update   *Order
	OnChange client.Script
}

type Order struct {
	clientside.String
}

func (order *Order) GetDefaultValue() client.Value {
	return client.NewString("[]")
}

//New makes the seed's children sortable and reorderable by a user.
func New(options ...seed.Option) seed.Option {

	return seed.Options{
		js.Require("/assets/js/sortable.js", ""),

		seed.Options(options),

		seed.NewOption(func(c seed.Seed) {
			var data data
			c.Load(&data)

			var OnSort client.Script

			if data.Update != nil {
				OnSort = data.Update.SetTo(js.String{Value: js.Func("JSON.stringify").Call(html.Element(c).Get("order"))})
			}

			OnSort = client.NewScript(
				OnSort,
				data.OnChange,
			)

			client.OnLoad(js.Script(func(q js.Ctx) {
				fmt.Fprintf(q, `%[1]v.sortable = new Sortable(%[1]v, {
					filter: ".sortable-ignore",
					onEnd: function(evt) {
						%[1]v.order = %[1]v.sortable.toArray();
						%[2]v();
					},
				}); %[1]v.order = %[1]v.sortable.toArray();`, html.Element(c), OnSort.GetValue().String())
			})).AddTo(c)

			change.On(client.If(html.Element(c).Get("sortable"),
				client.After(1, html.Element(c).Get("sortable").Run("sort", html.Element(c).Get("order"))),
			)).AddTo(c)
		}),
	}
}

//SetIDTo sets the sortable ID used for identifying an element.
func SetIDTo(s client.String) seed.Option {
	return attr.SetTo("data-id", s)
}

//When enables sorting by the user only when the provided condition is true.
func When(condition client.Bool) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(condition, c)

		c.With(
			change.On(
				html.Element(c).Get("sortable").Run("option", client.NewString("disabled"), not.True(condition)),
			),
		)
	})
}

//OnChange runs the provided script when the order of the sortable is changed by the user.
func OnChange(do client.Script) seed.Option {
	return seed.Mutate(func(data *data) {
		data.OnChange = do
	})
}

//Update sorts the array in the order specified by the given string.
func Update(variable *Order) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)

		c.With(
			variable.OnChange(
				html.Element(c).Set("order", js.Func("JSON.parse").Call(variable)),
			),

			seed.Mutate(func(data *data) {
				data.Update = variable
			}),
		)
	})
}
