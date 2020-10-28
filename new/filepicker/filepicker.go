package filepicker

import (
	"encoding/base64"
	"fmt"
	"math/big"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"

	"qlova.org/seed/new/html/input"
)

//New returns a new filepicker widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(attr.Set("type", "file"), seed.Options(options))
}

//Update the attachment when a file is picked.
func Update(a File) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		c.With(
			client.OnInput(js.Script(func(q js.Ctx) {
				q(js.Global().Get("seed").Set("active", html.Element(c)))
				q("try {")
				q(a.Set(js.NewValue(`event.target.files[0]`)))
				q("} catch(e) { seed.report(e) }")
			})),
		)
	})
}

//SelectFile asks the user to select a file, if and only if they do select a file, the handler will be called.
//There is no way to tell if the user did not select a file.
func SelectFile(handler func(File) client.Script) js.Script {
	var a = NewFile()

	return func(q js.Ctx) {
		var input = js.Global().Get("document").Call("createElement", js.NewString("input")).Var(q)
		q(input.Set("type", js.NewString("file")))

		q("let active = seed.active;")
		q("console.log(active.id);")

		q(input.Set("onchange", js.NewFunction(func(q js.Ctx) {
			q(js.Global().Get("seed").Set("active", js.NewValue(`active`)))

			q("try {")
			q(a.Set(js.NewValue(`event.target.files[0]`)))
			q(handler(a))
			q("} catch(e) { seed.report(e) }")
		}, "event")))

		q(input.Run("click"))
	}
}

type File struct {
	string
}

func (a File) Set(v js.Value) func(js.Ctx) {
	return func(q js.Ctx) {
		q(`if (!window.attachments) window.attachments = {};`)
		q(fmt.Sprintf(`window.attachments["%v"] = %v;`, a.string, v))
	}
}

func (a File) GetFile() js.Value {
	return js.NewValue(`window.attachments[%v]`, js.NewString(a.string))
}

func (a File) GetValue() js.Value {
	return a.GetFile()
}

func (a File) GetBool() js.Bool {
	return a.GetValue().GetBool()
}

var id int64

func NewFile() File {
	id++
	return File{base64.RawURLEncoding.EncodeToString(big.NewInt(int64(id)).Bytes())}
}
