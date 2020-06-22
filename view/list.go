package view

import (
	"fmt"

	"qlova.org/seed"
	"qlova.org/seed/client/render"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
)

//List creates a viewset that can be used with Index, Back & Next functions.
func List(views ...View) seed.Option {
	if len(views) == 0 {
		return seed.NewOption(func(seed.Seed) {})
	}

	return seed.Options{
		Set(views[0]),

		seed.NewOption(func(c seed.Seed) {
			script.OnReady(func(q script.Ctx) {
				q(script.Element(c).Set("view", js.NewFunction(func(q script.Ctx) {
					q("if (" + script.Element(c).Get("view").Get("index").String() + " < 0)")
					q(script.Element(c).Get("view").Set("index", js.NewNumber(0)))
					q("if (" + script.Element(c).Get("view").Get("index").String() + " >= " + fmt.Sprint(len(views)) + ")")
					q(script.Element(c).Get("view").Set("index", js.NewNumber(float64(len(views)-1))))
					q("switch (" + script.Element(c).Get("view").Get("index").String() + ") {")
					for i, v := range views {
						fmt.Fprintf(q, "case %v:", i)
						q(ControllerOf(c).Goto(v))
						q("return;")
					}
					q("}")
				})))
				q(script.Element(c).Get("view").Set("index", js.NewNumber(0)))
			}).AddTo(c)

			render.On(script.Element(c).Run("view", script.Element(c).Get("view").Get("index"))).AddTo(c)
		}),
	}
}

//Next changes to the next view in the List
func (c Controller) Next() js.Script {
	return script.New(
		script.Element(c.of).Get("view").Set("index", js.Number{script.Element(c.of).Get("view").Get("index")}.Plus(js.NewNumber(1))),
		script.Element(c.of).Run("view", script.Element(c.of).Get("view").Get("index")),
	)
}

//Back changes to the back view in the List
func (c Controller) Back() js.Script {
	return script.New(
		script.Element(c.of).Get("view").Set("index", js.Number{script.Element(c.of).Get("view").Get("index")}.Minus(js.NewNumber(1))),
		script.Element(c.of).Run("view", script.Element(c.of).Get("view.index")),
	)
}
