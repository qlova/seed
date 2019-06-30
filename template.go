package seed

type Template struct {
	Seed
}

func NewTemplate() Template {
	var seed = New()
	seed.template = true
	return Template{seed}
}

func (template Template) render(q Script) string {
	var seed = template.Seed

	q.Javascript("let "+seed.id+" = document.createElement(\"")
	q.Javascript(seed.tag)
	q.Javascript("\");")

	q.Javascript(seed.id+".className = '"+seed.id+"';")

	for _, child := range seed.children {
		q.Javascript(seed.id+".appendChild("+Template{child.Root()}.render(q)+");")
	}

	return seed.id
}

func (template Template) scripts(q Script) {
	var seed = template.Seed

	if seed.onready != nil {
		q.Javascript("{")
		seed.onready(q)
		q.Javascript("};")
	}
}

func (template Template) Render(q Script) string {
	var element = template.render(q)
	template.scripts(q)
	return element
}
