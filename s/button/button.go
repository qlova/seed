package button

import (
	"qlova.org/seed"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/state"
)

//New returns a new text widget.
func New(text string, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("button"),
		html.SetInnerText(text),
		attr.Set("type", "button"),
		seed.Options(options),
	)
}

func Var(text state.String, options ...seed.Option) seed.Seed {
	return New("", text.SetText(), seed.Options(options))
}
