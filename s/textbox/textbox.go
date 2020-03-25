package textbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
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
			text.Set(q.Value(`%v.value`, q.Scope(c).Element()).String())(q)
		}))
	}).And(options...))
}

//SetPlaceholder sets the placeholder of the textbox.
func SetPlaceholder(placeholder string) seed.Option {
	return attr.Set("placeholder", placeholder)
}
