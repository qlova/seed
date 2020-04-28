package textbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"

	"github.com/qlova/seed/s/html/input"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(options...)
}

//Var returns text with a variable text argument.
func Var(text state.String, options ...seed.Option) seed.Seed {
	return New(seed.Do(func(c seed.Seed) {
		c.Add(script.On("input", func(q script.Ctx) {
			text.Set(js.String{js.NewValue(script.Scope(c, q).Element() + `.value`)})(q)
		}), text.SetValue())
	}), seed.Options(options))
}

//SetPlaceholder sets the placeholder of the textbox.
func SetPlaceholder(placeholder string) seed.Option {
	return attr.Set("placeholder", placeholder)
}

//SetReadOnly sets the textbox to be readonly.
func SetReadOnly() seed.Option {
	return attr.Set("readonly", "")
}

//Focus focuses the textbox.
func Focus(c seed.Seed) script.Script {
	return func(q script.Ctx) {
		q(script.Element(c).Run(`focus`))
	}
}
