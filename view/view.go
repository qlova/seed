//Package view provides local view that can be swapped in and out.
package view

import (
	"reflect"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
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

//Goto returns a script that goes to the given view.
func (c Controller) Goto(view View) js.Script {
	//Sort out script arguments of the page.
	view, args := parseArgs(view)

	var data data
	c.of.Read(&data)

	var key = reflect.TypeOf(view)
	if _, ok := data.views[key]; !ok {
		if data.views == nil {
			data.views = make(map[reflect.Type]bool)
			c.of.Write(data)
		}

		data.views[reflect.TypeOf(view)] = true

		var template = seed.New(
			html.SetTag("template"),
		)
		template.Use()
		template.AddTo(c.of)

		var element = view.View(c)
		element.Add(
			html.AddClass(Name(view)),
		)
		element.Use()
		element.AddTo(template)
	}

	c.of.Write(data)

	return func(q script.Ctx) {
		q.Run(`seed.view`, js.NewValue(script.Scope(c.of, q).Element()), js.NewString(Name(view)), args)
	}
}

//View is a local view.
type View interface {
	View(Controller) Seed
}

type Seed struct {
	seed.Seed
}

func New(options ...seed.Option) Seed {
	return Seed{seed.New(
		html.SetTag("div"),

		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),

		seed.Do(func(c seed.Seed) {
			c.Add(
				OnEnter(state.Refresh(c)),
			)
		}),

		seed.Options(options),
	)}
}

type data struct {
	seed.Data

	views map[reflect.Type]bool
}

func OnEnter(f script.Script) seed.Option {
	return script.On("viewenter", f)
}

func OnExit(f script.Script) seed.Option {
	return script.On("viewexit", f)
}

func State(v View) state.State {
	return state.New(state.SetKey("view."+Name(v)), state.ReadOnly())
}
