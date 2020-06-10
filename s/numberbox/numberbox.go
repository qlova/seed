package numberbox

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
	return input.New(attr.Set("type", "number"), seed.Options(options))
}

//Var returns text with a variable text argument.
func Var(f state.Float, options ...seed.Option) seed.Seed {
	if f.Null() {
		return New(options...)
	}
	return New(seed.NewOption(func(c seed.Seed) {
		c.With(script.On("input", func(q script.Ctx) {
			f.Set(js.Number{js.NewValue(script.Scope(c, q).Element() + `.value`)})(q)
		}), state.SetProperty("value", f))
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
