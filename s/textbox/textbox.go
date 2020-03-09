package textbox

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/state"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("input").And(options...),
	)
}

//Var returns text with a variable text argument.
func Var(text state.String, options ...seed.Option) seed.Seed {
	return New(seed.Do(func(c seed.Seed) {
		c.Add(script.On("input", func(q script.Ctx) {
			text.Set(q.Value(`%v.value`, c.Ctx(q).Element()).String())(q)
		}))
	}).And(options...))
}
