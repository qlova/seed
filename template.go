package seed

import "github.com/qlova/seed/style/css"

type Template struct {
	Seed
}

func NewTemplate() Template {
	var seed = New()
	seed.template = true

	seed.SetDisplay(css.Flex)
	seed.SetFlexDirection(css.Column)
	seed.SetFlexWrap(css.Wrap)

	return Template{seed}
}

func (template Template) render(q Script) string {
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

func (seed Seed) Template() {
	seed.template = true
}

func (template Template) scripts(q Script) {
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

func (template Template) Render(q Script) string {
	var element = template.render(q)
	template.scripts(q)
	return element
}
