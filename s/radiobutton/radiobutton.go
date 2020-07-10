package radiobutton

import (
	"qlova.org/seed"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/script"

	"qlova.org/seed/s/html/input"
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
func Focus(c seed.Seed) script.Script {
	return func(q script.Ctx) {
		q(script.Element(c).Run(`focus`))
	}
}
