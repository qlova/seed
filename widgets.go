package seed

func ToolBar() Seed {
	return NewToolBar()
}

func NewToolBar() Seed {
	seed := New()
	seed.SetLayout("flex")
	seed.SetSticky()
	return seed
}


func Text() Seed {
	seed := New()
	seed.tag = "p"
	return seed
}


func LolScript(content string) Seed {
	seed := New()
	seed.tag = "script"
	seed.content = []byte(content)
	return seed
}

func Header() Seed {
	seed := New()
	seed.tag = "h1"
	return seed
}

func FilePicker(types string) Seed {
	seed := New()
	seed.tag = "input"
	seed.attr = `type="file" accept="`+types+`"`
	return seed
}

func TextBox() Seed {
	seed := New()
	seed.tag = "input"
	return seed
}

func TextArea() Seed {
	seed := New()
	seed.tag = "textarea"
	return seed
}

func Button() Seed {
	seed := New()
	seed.tag = "button"
	return seed
}
