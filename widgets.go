package seed

import "fmt"
import "image/color"

import "github.com/qlova/seed/style/css"

func NewToolBar() Seed {
	seed := New()

	seed.SetName("Toolbar")
	seed.Stylable.Set("display", "flex")
	seed.Stylable.Set("position", "fixed")
	seed.SetFlexDirection(css.Row)

	return seed
}

func NewSpacer(amount ...float64) Seed {
	
	
	seed := New()
	seed.SetName("Spacer")

	if len(amount) > 0 {
		seed.SetExpand(amount[0])
	} else {
		seed.SetExpand(1)
	}
	
	return seed
}

func NewExpander(ratio ...float64) Seed {
	
	seed := New()
	if len(ratio) > 0 {
		seed.SetExpand(ratio[0])
	} else {
		seed.SetExpand(1)
	}
	
	return seed
}

func AddExpanderTo(parent Seed, ratio ...float64) Seed {
	seed := NewExpander(ratio...)
	parent.Add(seed)
	return seed
}


func NewLine() Seed {
	seed := New()
	seed.SetName("Line")
	seed.tag = "hr"
	
	seed.SetSize(Auto, Auto)

	seed.Set("border-style", "solid")
	
	return seed
}

func NewLink(url string) Seed {
	seed := New()
	seed.SetName("Link")
	seed.tag = "a"
	seed.attr = "href='"+url+"'"
	
	return seed
}

func NewHeader() Seed {
	seed := New()
	seed.SetName("Header")
	seed.tag = "h1"
	
	seed.SetSize(Auto, Auto)
	
	return seed
}

func NewFilePicker(types string) Seed {
	seed := New()
	seed.SetName("File")
	seed.tag = "input"
	seed.attr = `type="file" accept="`+types+`"`
	return seed
}

func NewTextArea() Seed {
	seed := New()
	seed.SetName("TextArea")
	seed.tag = "textarea"
	seed.attr = "data-gramm_editor=false"
	
	return seed
}

func NewListBox(values []string) Seed {
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

//A widget that displays text.
type Text struct {
	Seed
}

//Set the text color.
func (text Text) SetColor(c color.Color) {
	text.SetTextColor(c)
}

//Set the text color.
func (text Text) SetSize(s complex128) {
	text.SetTextSize(s)
}



func NewText(s ...string) Text {
	seed := New()
	seed.SetName("Text")
	seed.tag = "span"
	
	if len(s) > 0 {
		seed.SetText(s[0])
	}
	
	seed.SetSize(Auto, Auto)
	
	var Text = Text{
		Seed: seed,
	}
	
	return Text
}


//Create a new Text widget and add it to the provided parent.
func AddTextTo(parent Interface, s ...string) Text {
	var Text = NewText(s...)
	parent.GetSeed().Add(Text)
	return Text
}

//A widget that displays text.
type TextBox struct {
	Seed
}

func NewTextBox(s ...string) TextBox {
	seed := New()
	seed.SetName("Text")
	seed.tag = "input"
	
	if len(s) > 0 {
		seed.SetText(s[0])
	}
	
	seed.SetSize(Auto, Auto)
	
	var TextBox = TextBox{
		Seed: seed,
	}
	return TextBox
}

//Create a new Text widget and add it to the provided parent.
func AddTextBoxTo(parent Interface, s ...string) TextBox {
	var TextBox = NewTextBox(s...)
	parent.GetSeed().Add(TextBox)
	return TextBox
}

//A widget that displays text.
type Space struct {
	Seed
}

//Create a new Text widget and add it to the provided parent.
func AddSpaceTo(parent Interface, s ...complex128) Space {
	seed := New()
	seed.SetName("Text")
	seed.tag = "div"
	
	if len(s) > 0 {
		seed.SetSize(s[0], s[0])
	}
	
	var Space = Space{
		Seed: seed,
	}
	parent.GetSeed().Add(Space)
	return Space
}

//A widget that displays text.
type PasswordBox struct {
	Seed
}

//Create a new Text widget and add it to the provided parent.
func AddPasswordBoxTo(parent Interface) PasswordBox {
	seed := New()
	seed.SetName("Text")
	seed.tag = "input"
	seed.attr = `type="password"`
	
	seed.SetSize(Auto, Auto)
	
	var PasswordBox = PasswordBox{
		Seed: seed,
	}
	parent.GetSeed().Add(PasswordBox)
	return PasswordBox
}

func NewButton() Seed {
	seed := New()
	seed.SetName("Button")
	seed.tag = "button"
	
	seed.SetSize(Auto, Auto)
	
	return seed
}

func AddButtonTo(parent Seed) Seed {
	seed := NewButton()
	parent.Add(seed)
	return seed
}

func NewRow() Seed {
	seed := New()
	seed.tag = "div"
	seed.SetName("Row")
	seed.Stylable.Set("display", "flex")
	seed.Stylable.Set("flex-direction", "row")
	seed.Stylable.Set("flex-shrink", "1")

	return seed
}

func AddRowTo(seed Seed) Seed {
	var row = NewRow()
	seed.Add(row)
	return row
}

func NewColumn() Seed {
	seed := New()
	seed.tag = "div"
	seed.SetName("Column")
	seed.Stylable.Set("display", "inline-flex")
	seed.Stylable.Set("flex-direction", "column")
	seed.Stylable.Set("flex-shrink", "1")


	return seed
}

func AddColumnTo(seed Seed) Seed {
	var column = NewColumn()
	seed.Add(column)
	return column
}