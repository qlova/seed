package seed

import "fmt"

func ToolBar() Seed {
	return NewToolBar()
}

func NewToolBar() Seed {
	seed := New()
	seed.Stylable.Set("display", "flex")
	seed.Stylable.Set("position", "fixed")
	return seed
}

func Line() Seed {
	seed := New()
	seed.tag = "hr"

	seed.Set("border-style", "solid")
	
	return seed
}

func Row() Seed {
	seed := New()
	seed.tag = "div"
	seed.Stylable.Set("display", "flex")
	seed.Stylable.Set("flex-direction", "row")
	seed.Stylable.Set("align-items", "center")
	return seed
}

func Col() Seed {
	seed := New()
	seed.tag = "div"
	seed.Stylable.Set("display", "inline-flex")
	seed.Stylable.Set("flex-direction", "column")
	seed.Stylable.Set("align-items", "center")
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
	seed.attr = "data-gramm_editor=false"
	return seed
}

func Button() Seed {
	seed := New()
	seed.tag = "button"
	return seed
}

func ListBox(values []string) Seed {
	seed := New()
	seed.tag = "select"
	
	var content string
	
	for _, value := range values {
		content += fmt.Sprint("<option value='", value, "'>", value, "</option>")
	}
	
	seed.SetContent(content)

	return seed
}
