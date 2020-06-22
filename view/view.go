//Package view provides local view that can be swapped in and out.
package view

import (
	"fmt"
	"reflect"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/style"
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

func (c Controller) State(v View) state.State {
	return state.New(state.SetKey(script.ID(c.of)+".view."+Name(v)),
		state.SetRaw(fmt.Sprintf("(q.get('%[1]v').CurrentView && q.get('%[1]v').CurrentView.classList.contains('%[2]v'))", script.ID(c.of), Name(v))),
		state.ReadOnly(),
	)
}

//Goto returns a script that goes to the given view.
func (c Controller) Goto(view View) js.Script {
	//Sort out script arguments of the page.
	view, args := parseArgs(view, c.of)

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
		element.With(
			html.AddClass(Name(view)),
		)
		element.Use()
		element.AddTo(template)
	}

	c.of.Write(data)

	var seed_view = js.Function{js.NewValue(`seed.view`)}

	return func(q script.Ctx) {
		q.Run(seed_view, js.NewValue(script.Scope(c.of, q).Element()), js.NewString(Name(view)), args)
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

		style.SetSize(100, 100),

		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),

		seed.NewOption(func(c seed.Seed) {
			c.With(
				OnEnter(clientside.Render(c)),
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
