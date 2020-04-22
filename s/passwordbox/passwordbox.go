package passwordbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state/secret"

	"github.com/qlova/seed/s/textbox"
)

//New returns a new passwordbox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "password").And(options...))
}

//Var returns an passwordbox with a variable state argument that will be synced with the value of this passwordbox.
func Var(text secret.State, options ...seed.Option) seed.Seed {
	return New(seed.Do(func(c seed.Seed) {
		c.Add(script.On("input", func(q script.Ctx) {
			text.Set(js.String{Value: js.NewValue(script.Scope(c, q).Element() + `.value`)})(q)
		}))
	}).And(options...))
}
