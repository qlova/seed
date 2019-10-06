package seed

import (
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style/css"
)

//Template is a seed that can be duplicated and modified based on data.
type Template struct {
	Seed
}

//NewTemplate returns a new template.
func NewTemplate() Template {
	var seed = New()
	seed.template = true

	seed.SetDisplay(css.Flex)
	seed.SetFlexDirection(css.Column)
	seed.SetFlexWrap(css.Wrap)

	return Template{seed}
}

func (template Template) render(q script.Ctx) string {
	var seed = template.Seed

	q.Javascript("let " + seed.id + " = document.createElement(\"")
	q.Javascript(seed.tag)
	q.Javascript("\");")

	if seed.content != nil {
		q.Javascript(seed.id + ".innerHTML = '" + string(seed.content) + "';")
	}

	q.Javascript(seed.id + ".className = '" + seed.id + "';")

	for attribute, value := range seed.Element.Attributes {
		q.Javascript(seed.id + ".setAttribute('" + string(attribute) + "', '" + value + "');")
	}

	for _, child := range seed.children {
		q.Javascript(seed.id + ".appendChild(" + Template{child.Root()}.render(q) + ");")
	}

	return seed.id
}

//Template returns true when the seed is a template.
func (seed Seed) Template() {
	seed.template = true
}

func (template Template) scripts(q script.Ctx) {
	var seed = template.Seed

	if seed.onready != nil {
		q.Javascript("{")
		seed.onready(q)
		q.Javascript("};")
	}

	for _, child := range seed.children {
		var template = Template{child.Root()}
		template.template = true
		template.scripts(q)
	}
}

//Render renders the template.
func (template Template) Render(q script.Ctx) string {
	var element = template.render(q)
	template.scripts(q)
	return element
}
