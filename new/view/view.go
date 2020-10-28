//Package view provides local view that can be swapped in and out.
package view

import (
	"reflect"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/web/css"
	"qlova.org/seed/web/html"
	"qlova.org/seed/web/js"
	"qlova.org/seed/set"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/web/css/units/percentage/of"
)

func Name(view View) string {
	return strings.Replace(reflect.TypeOf(view).String(), ".", "_", -1)
}

//Controller is responsible for showing the current view.
type Controller struct {
	of seed.Seed
}

//ControllerOf returns the Controller for the given seed.
func ControllerOf(c seed.Seed) Controller {
	return Controller{c}
}

func (c Controller) Seed() seed.Seed {
	return c.of
}

//Is returns true if the given view is the current view of the given controller..
func (c Controller) Is(v View) *clientside.Bool {
	return &clientside.Bool{
		Name: client.ID(c.of) + ".view." + Name(v),
	}
}

//Goto returns a script that goes to the given view.
func (c Controller) Goto(view View) js.Script {
	//Sort out script arguments of the page.
	view, args := parseArgs(view, c.of)

	var data data
	c.of.Load(&data)

	var key = reflect.TypeOf(view)
	if _, ok := data.views[key]; !ok {
		if data.views == nil {
			data.views = make(map[reflect.Type]bool)
			c.of.Save(data)
		}

		data.views[reflect.TypeOf(view)] = true

		var template = seed.New(
			html.SetTag("template"),
		)
		template.Use()
		template.AddTo(c.of)

		var element = view.View(c)
		element.With(
			html.AddClass(Name(view)),
		)
		element.Use()
		element.AddTo(template)
	}

	c.of.Save(data)

	var seed_view = js.Function{js.NewValue(`seed.view`)}

	return func(q js.Ctx) {
		q.Run(seed_view, html.Element(c.of), js.NewString(Name(view)), args)
	}
}

//View is a local view.
type View interface {
	View(Controller) seed.Seed
}

func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("div"),

		set.Size(100%of.Parent, 100%of.Parent),

		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),

		seed.NewOption(func(c seed.Seed) {
			c.With(
				OnEnter(clientside.Render(c)),
			)
		}),

		transition.SetOnEnter(OnEnter),
		transition.SetOnExit(OnExit),

		seed.Options(options),
	)
}

type data struct {
	views map[reflect.Type]bool
}

func OnEnter(f ...client.Script) seed.Option {
	return client.On("viewenter", f...)
}

func OnExit(f ...client.Script) seed.Option {
	return client.On("viewexit", f...)
}
