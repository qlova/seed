//Package page provides global pages that can be swapped in and out.
package page

import (
	"reflect"
	"strings"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/style"
)

//ID returns the DOM id of the provided page.
func ID(p Page) string {
	return "." + strings.Replace(reflect.TypeOf(p).String(), ".", "_", -1)
}

//Router is responsible for showing the current page and routing urls to the approproate page.
type Router struct {
	c seed.Seed
}

//RouterOf returns the Router for the given seed.
func RouterOf(c seed.Seed) Router {
	return Router{c}
}

//Goto returns a script that goes to the given page.
func (r Router) Goto(page Page) js.Script {
	return func(q script.Ctx) {
		//Sort out script arguments of the page.
		page, args, path := parseArgs(page)

		var data data
		r.c.Read(&data)
		data.pages = append(data.pages, page)
		r.c.Write(data)

		q.Run(`seed.goto`, js.NewString(ID(page)), args, path)
	}

}

//Page is a global view.
type Page interface {
	Page(Router) Seed
}

type Seed struct {
	seed.Seed
}

func New(options ...seed.Option) Seed {
	return Seed{seed.New(
		html.SetTag("div"),
		html.AddClass("page"),

		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),
		style.Expand(),

		style.SetWidth(100),

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

	pages []Page
}

var seeds = make(map[seed.Seed]data)

func OnEnter(f script.Script) seed.Option {
	return script.On("pageenter", f)
}

func OnExit(f script.Script) seed.Option {
	return script.On("pageexit", f)
}

func State(p Page) state.State {
	return state.New(state.SetKey("page."+ID(p)), state.ReadOnly())
}
