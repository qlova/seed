package text

import "image/color"
import "github.com/qlova/seed"

type Widget struct {
	seed.Seed
}

//Set the text color.
func (text Widget) SetColor(c color.Color) {
	text.SetTextColor(c)
}

//Set the text's font-size.
func (text Widget) SetSize(s complex128) {
	text.SetTextSize(s)
}

func New(s ...string) Widget {
	widget := seed.New()
	widget.SetTag("span")
	
	if len(s) > 0 {
		widget.SetText(s[0])
	}
	
	widget.SetSize(seed.Auto, seed.Auto)

	widget.Align(0)

	return  Widget{widget}
}

//Create a new Text widget and add it to the provided parent.
func AddTo(parent seed.Interface, s ...string) Widget {
	var Text = New(s...)
	parent.Root().Add(Text)
	return Text
}