package filepicker

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html/attr"
	"github.com/qlova/seed/js"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/user/attachment"

	"github.com/qlova/seed/s/html/input"
)

//File is a type of attachment.
type File = attachment.Attachment

//New returns a new filepicker widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(attr.Set("type", "file"), seed.Options(options))
}

//Var returns a filepicker that sets the attachment to reflect the user's input.
func Var(a attachment.Attachment, options ...seed.Option) seed.Seed {
	return New(seed.NewOption(func(c seed.Seed) {
		c.With(script.On("input", func(q script.Ctx) {
			a.Set(js.NewValue(script.Scope(c, q).Element() + `.files[0]`))(q)
		}))
	}), seed.Options(options))
}

//SelectFile asks the user to select a file, if and only if they do select a file, the handler will be called.
//There is no way to tell if the user did not select a file.
func SelectFile(handler func(File) script.Script) script.Script {
	var a = attachment.New()

	return func(q script.Ctx) {
		var input = js.Global().Get("document").Call("createElement", js.NewString("input")).Var(q)
		q(input.Set("type", js.NewString("file")))

		q("let active = seed.active;")
		q("console.log(active.id);")

		q(input.Set("onchange", js.NewFunction(func(q script.Ctx) {
			q(js.Global().Get("seed").Set("active", js.NewValue(`active`)))

			q("try {")
			q(a.Set(js.NewValue(`arguments[0].target.files[0]`)))
			q(handler(a))
			q("} catch(e) { seed.report(e) }")
		})))

		q(input.Run("click"))
	}
}
