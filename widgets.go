package seed

import "fmt"

func ToolBar() Seed {
	return NewToolBar()
}

func NewToolBar() Seed {
	seed := New()
	seed.SetName("Toolbar")
	seed.Stylable.Set("display", "flex")
	seed.Stylable.Set("position", "fixed")
	return seed
}

func Line() Seed {
	seed := New()
	seed.SetName("Line")
	seed.tag = "hr"

	seed.Set("border-style", "solid")
	
	return seed
}

func Row() Seed {
	seed := New()
	seed.tag = "div"
	seed.SetName("Row")
	seed.Stylable.Set("display", "flex")
	seed.Stylable.Set("flex-direction", "row")
	return seed
}

func Col() Seed {
	seed := New()
	seed.tag = "div"
	seed.SetName("Column")
	seed.Stylable.Set("display", "inline-flex")
	seed.Stylable.Set("flex-direction", "column")
	return seed
}

func Text() Seed {
	seed := New()
	seed.SetName("Text")
	seed.tag = "p"
	return seed
}

func Header() Seed {
	seed := New()
	seed.SetName("Header")
	seed.tag = "h1"
	return seed
}

func FilePicker(types string) Seed {
	seed := New()
	seed.SetName("File")
	seed.tag = "input"
	seed.attr = `type="file" accept="`+types+`"`
	return seed
}

func TextBox() Seed {
	seed := New()
	seed.SetName("TextBox")
	seed.tag = "input"
	return seed
}

func TextArea() Seed {
	seed := New()
	seed.SetName("TextArea")
	seed.tag = "textarea"
	seed.attr = "data-gramm_editor=false"
	return seed
}

func Button() Seed {
	seed := New()
	seed.SetName("Button")
	seed.tag = "button"
	return seed
}

func ListBox(values []string) Seed {
	seed := New()
	seed.SetName("ListBox")
	seed.tag = "select"
	
	var content string
	
	for _, value := range values {
		content += fmt.Sprint("<option value='", value, "'>", value, "</option>")
	}
	
	seed.SetContent(content)

	return seed
}
