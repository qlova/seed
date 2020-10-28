package textarea

import (
	"qlova.org/seed"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/web/css"
	"qlova.org/seed/web/html"
	"qlova.org/seed/web/html/attr"
	"qlova.org/seed/web/js"

	"qlova.org/seed/new/html/textarea"
	"qlova.org/seed/new/textbox"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return textarea.New(
		css.SetResize(css.None),
		seed.Options(options),
	)
}

//Update updates the given variable whenever the textbox text is modified.
func Update(variable *clientside.String) seed.Option {
	return textbox.Update(variable)
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
func Focus(c seed.Seed) js.Script {
	return func(q js.Ctx) {
		q(html.Element(c).Run(`focus`))
	}
}
