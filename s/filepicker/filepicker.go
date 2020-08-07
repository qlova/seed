package filepicker

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/user/attachment"

	"qlova.org/seed/s/html/input"
)

//File is a type of attachment.
type File = attachment.Attachment

//New returns a new filepicker widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(attr.Set("type", "file"), seed.Options(options))
}

//Update the attachment when a file is picked.
func Update(a attachment.Attachment) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		c.With(
			client.OnInput(js.Script(func(q script.Ctx) {
				q(js.Global().Get("seed").Set("active", script.Element(c)))
				q("try {")
				q(a.Set(js.NewValue(`event.target.files[0]`)))
				q("} catch(e) { seed.report(e) }")
			})),
		)
	})
}

//SelectFile asks the user to select a file, if and only if they do select a file, the handler will be called.
//There is no way to tell if the user did not select a file.
func SelectFile(handler func(File) client.Script) script.Script {
	var a = attachment.New()

	return func(q script.Ctx) {
		var input = js.Global().Get("document").Call("createElement", js.NewString("input")).Var(q)
		q(input.Set("type", js.NewString("file")))

		q("let active = seed.active;")
		q("console.log(active.id);")

		q(input.Set("onchange", js.NewFunction(func(q script.Ctx) {
			q(js.Global().Get("seed").Set("active", js.NewValue(`active`)))

			q("try {")
			q(a.Set(js.NewValue(`event.target.files[0]`)))
			q(handler(a))
			q("} catch(e) { seed.report(e) }")
		}, "event")))

		q(input.Run("click"))
	}
}
