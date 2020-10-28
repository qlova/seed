package view

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/set/change"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/js"
)

//List creates a viewset that can be used with Index, Back & Next functions.
func List(views ...View) seed.Option {
	if len(views) == 0 {
		return seed.NewOption(func(seed.Seed) {})
	}

	return seed.Options{
		Set(views[0]),

		seed.NewOption(func(c seed.Seed) {
			client.OnLoad(js.Script(func(q js.Ctx) {
				q(html.Element(c).Set("view", js.NewFunction(func(q js.Ctx) {
					q("if (" + html.Element(c).Get("view").Get("index").String() + " < 0)")
					q(html.Element(c).Get("view").Set("index", js.NewNumber(0)))
					q("if (" + html.Element(c).Get("view").Get("index").String() + " >= " + fmt.Sprint(len(views)) + ")")
					q(html.Element(c).Get("view").Set("index", js.NewNumber(float64(len(views)-1))))
					q("switch (" + html.Element(c).Get("view").Get("index").String() + ") {")
					for i, v := range views {
						fmt.Fprintf(q, "case %v:", i)
						q(ControllerOf(c).Goto(v))
						q("return;")
					}
					q("}")
				})))
				q(html.Element(c).Get("view").Set("index", js.NewNumber(0)))
			})).AddTo(c)

			change.On(html.Element(c).Run("view", html.Element(c).Get("view").Get("index"))).AddTo(c)
		}),
	}
}

//Reset the controller to the default view.
func (c Controller) Reset() client.Script {
	return client.NewScript(
		html.Element(c.of).Get("view").Set("index", js.NewNumber(0)),
		html.Element(c.of).Run("view", html.Element(c.of).Get("view").Get("index")),
	)
}

//Next changes to the next view in the List
func (c Controller) Next() client.Script {
	return client.NewScript(
		html.Element(c.of).Get("view").Set("index", js.Number{html.Element(c.of).Get("view").Get("index")}.Plus(js.NewNumber(1))),
		html.Element(c.of).Run("view", html.Element(c.of).Get("view").Get("index")),
	)
}

//Back changes to the back view in the List
func (c Controller) Back() client.Script {
	return client.NewScript(
		html.Element(c.of).Get("view").Set("index", js.Number{html.Element(c.of).Get("view").Get("index")}.Minus(js.NewNumber(1))),
		html.Element(c.of).Run("view", html.Element(c.of).Get("view.index")),
	)
}
