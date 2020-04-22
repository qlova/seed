package text

import (
	"strings"

	html_go "html"

	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/state"
)

//New returns a new text widget.
func New(text state.AnyString, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("p"),
		state.SetText(text),
		seed.Options(options),
	)
}

//Var returns text with a variable text argument.
func Var(text state.String, options ...seed.Option) seed.Seed {
	return New("", text.SetText().And(options...))
}

//Set sets the text content of the text.
func Set(value string) seed.Option {
	value = html_go.EscapeString(value)
	value = strings.Replace(value, "\n", "<br>", -1)
	value = strings.Replace(value, "  ", "&nbsp;&nbsp;", -1)
	value = strings.Replace(value, "\t", "&emsp;", -1)

	return html.Set(value)
}
