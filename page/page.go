//Package page provides global pages that can be swapped in and out.
package page

import (
	"reflect"
	"strconv"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/style"
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

		q.Run(js.Function{js.NewValue(`seed.goto`)}, js.NewString(ID(page)), args, path)
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
	var Page = Seed{seed.New()}

	Page.With(
		html.SetTag("div"),

		css.SetDisplay(css.Flex),
		css.SetFlexDirection(css.Column),
		style.Expand(),

		style.SetSize(100, 100),
	)

	for _, option := range options {
		option.AddTo(Page)
	}

	return Page
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
	return state.New(state.SetKey("page"+ID(p)), state.SetRaw("(seed.CurrentPage && seed.CurrentPage.classList.contains("+strconv.Quote(ID(p)[1:])+"))"), state.ReadOnly())
}

type EnterIfOption struct {
	condition js.AnyBool
	otherwise script.Script
}

func EnterIf(condition js.AnyBool) EnterIfOption {
	return EnterIfOption{condition, nil}
}

func (e EnterIfOption) Else(do script.Script) seed.Option {
	e.otherwise = do
	return e
}

func (e EnterIfOption) AddTo(c seed.Seed) {
	var condition = js.NewObject{
		"test": js.NewFunction(js.Return(e.condition)),
	}

	if e.otherwise != nil {
		condition["otherwise"] = js.NewFunction(e.otherwise)
	}

	c.With(
		script.OnReady(script.New(
			script.Element(c).Set("conditions", js.NewValue(`(%v || [])`, script.Element(c).Get("conditions"))),
			script.Element(c).Get("conditions").Run("push", condition),
		)),
	)
}
