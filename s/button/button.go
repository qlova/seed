package button

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/state"
)

//New returns a new text widget.
func New(text string, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("button"),
		html.SetInnerText(text),
		seed.Options(options),
	)
}

func Var(text state.String, options ...seed.Option) seed.Seed {
	return New("", text.SetText(), seed.Options(options))
}
