package filepicker

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user/attachment"

	"github.com/qlova/seed/s/html/input"
)

//New returns a new filepicker widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(attr.Set("type", "file").And(options...))
}

//Var returns a filepicker that sets the attachment to reflect the user's input.
func Var(a attachment.Attachment, options ...seed.Option) seed.Seed {
	return New(seed.Do(func(c seed.Seed) {
		c.Add(script.On("input", func(q script.Ctx) {
			a.Set(js.NewValue(script.Scope(c, q).Element() + `.files[0]`))(q)
		}))
	}).And(options...))
}
