//Package page provides global pages that can be swapped in and out.
package page

import (
	"reflect"
	"strings"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/js"
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
	return func(q js.Ctx) {
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

func OnEnter(f client.Script) seed.Option {
	return client.On("pageenter", f)
}

func OnExit(f client.Script) seed.Option {
	return client.On("pageexit", f)
}

//Is returns true if the given page is the current page.
func Is(p Page) *clientside.Bool {
	return &clientside.Bool{
		Name: "page " + ID(p)[1:],
	}
}

type EnterIfOption struct {
	condition js.AnyBool
	otherwise client.Script
}

func EnterIf(condition js.AnyBool) EnterIfOption {
	return EnterIfOption{condition, nil}
}

func (e EnterIfOption) Else(do client.Script) seed.Option {
	e.otherwise = do
	return e
}

func (e EnterIfOption) AddTo(c seed.Seed) {
	var condition = js.NewObject{
		"test": js.NewFunction(js.Return(e.condition)),
	}

	if e.otherwise != nil {
		condition["otherwise"] = e.otherwise.GetScript().GetFunction()
	}

	c.With(
		client.OnLoad(client.NewScript(
			html.Element(c).Set("conditions", js.NewValue(`(%v || [])`, html.Element(c).Get("conditions"))),
			html.Element(c).Get("conditions").Run("push", condition),
		)),
	)
}
