package template

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/unit"
)

//Seed is a template that can be duplicated and modified based on provided data.
type Seed struct {
	seed.Seed
	refresh *func(script.Ctx)
}

//New returns a new template.
func New() Seed {
	var seed = Seed{seed.New(), new(func(script.Ctx))}
	seed.SetSize(100, unit.Auto)
	seed.SetUnshrinkable()

	seed.CSS().SetDisplay(css.Flex)
	seed.CSS().SetFlexDirection(css.Row)

	seed.OnReady(func(q script.Ctx) {
		var ctx = seed.Ctx(q)

		q.Javascript(`%v.refresh = function(feed) {`, ctx.Element())
		q.Javascript(`var %[2]v = %[1]v; %[2]v.innerHTML = ""; `, ctx.Element(), ctx.ID)

		q.Javascript(`if (!feed) return;`)

		q.Require(script.Get)
		q.Javascript(`if (%v.onrefresh) %[1]v.onrefresh();`, ctx.ID)

		q.Javascript(`if (!Array.isArray(feed)) feed = [feed];`)
		q.Javascript(`for (let data of feed) {`)
		{
			seed.renderTo(q, q.Value("data").Native())
			seed.renderScriptsTo(q, q.Value("data").Native())
		}
		q.Javascript(`}};`)
	})

	seed.TemplateRoot = true

	return seed
}

//AddTo parent.
func AddTo(parent seed.Interface) Seed {
	var template = New()
	parent.Root().Add(template)
	return template
}

//OnRefresh is called when the data of the template is refreshed.
func (template Seed) OnRefresh(f func(q script.Ctx, data script.Dynamic)) {
	var old = *template.refresh
	*template.refresh = func(q script.Ctx) {
		if old != nil {
			old(q)
		}
		f(q, template.Ctx(q).Data().Dynamic())
	}
}

func (template Seed) renderScriptsTo(q script.Ctx, data script.Type) {
	for _, child := range template.Children() {

		if ready := child.Root().Ready(); ready != nil {
			q.Javascript("{")
			ready(q)
			q.Javascript("};")
		}

		var template = Seed{child.Root(), new(func(script.Ctx))}
		template.Template = true
		template.renderScriptsTo(q, data)
	}

	if *template.refresh != nil {
		(*template.refresh)(q)
	}
}

func (template Seed) renderTo(q script.Ctx, data script.Type) string {
	for _, child := range template.Children() {
		var child = child.Root()

		q.Javascript("let " + child.ID() + " = document.createElement(\"")
		q.Javascript(child.Tag())
		q.Javascript("\");")

		if html := child.HTML(); html != nil {
			q.Javascript(child.ID() + ".innerHTML = '" + string(html) + "';")
		}

		q.Javascript(child.ID() + ".className = '" + child.ID() + "';")

		for attribute, value := range child.Element.Attributes {
			q.Javascript(child.ID() + ".setAttribute('" + string(attribute) + "', '" + value + "');")
		}

		var wrapped = Seed{child.Root(), new(func(script.Ctx))}
		wrapped.Template = true
		q.Javascript(template.ID() + ".appendChild(" + wrapped.renderTo(q, data) + ");")
	}

	return template.ID()
}

//Ctx is a template script context.
type Ctx struct {
	template Seed
	script.Seed
}

//Ctx returns the template's context.
func (template Seed) Ctx(q script.Ctx) Ctx {
	return Ctx{template, template.Seed.Ctx(q)}
}

//Data returns a reference to the template's data.
func (ctx Ctx) Data() script.Value {
	return ctx.Q.Value("data")
}

//Refresh refreshes the template with the provided data.
func (ctx Ctx) Refresh(data interface{}) {

	switch v := data.(type) {
	case script.Type:
		ctx.Q.Javascript("%v.refresh(%v);", ctx.Element(), v.LanguageType().Raw())
	case func(user seed.User):
		ctx.Q.Go(v).Then(func(d script.Dynamic) {
			ctx.Refresh(d)
		})
	default:
		panic("template: invalid data")
	}

}
