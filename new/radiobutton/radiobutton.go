package radiobutton

import (
	"qlova.org/seed"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/script"

	"qlova.org/seed/new/html/input"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(attr.Set("type", "radiobox"), seed.Options(options))
}

//SetReadOnly sets the textbox to be readonly.
func SetReadOnly() seed.Option {
	return attr.Set("readonly", "")
}

//Focus focuses the textbox.
func Focus(c seed.Seed) js.Script {
	return func(q js.Ctx) {
		q(html.Element(c).Run(`focus`))
	}
}
