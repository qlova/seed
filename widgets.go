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

func Row() Seed {
	seed := New()
	seed.tag = "div"
	seed.Style.Css.Set("display", "flex")
	seed.Style.Css.Set("flex-direction", "row")
	seed.Style.Css.Set("align-items", "center")
	return seed
}

func Col() Seed {
	seed := New()
	seed.tag = "div"
	seed.Style.Css.Set("display", "flex")
	seed.Style.Css.Set("flex-direction", "column")
	seed.Style.Css.Set("align-items", "center")
	return seed
}

func Text() Seed {
	seed := New()
	seed.tag = "p"
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
