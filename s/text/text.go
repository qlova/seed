package text

import (
	"image/color"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/state"
	"github.com/qlova/seed/style"
)

//New returns a new text widget.
func New(text string, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("p"),
		html.SetInnerText(text).And(options...),
	)
}

//Var returns text with a variable text argument.
func Var(text state.String, options ...seed.Option) seed.Seed {
	return New("", text.SetText().And(options...))
}

//SetColor sets the color of the text.
func SetColor(c color.Color) seed.Option {
	return style.SetTextColor(c)
}

//Set sets the text content of the text.
func Set(value string) seed.Option {
	return html.SetInnerText(value)
}
