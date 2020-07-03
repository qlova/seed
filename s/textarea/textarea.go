package textarea

import (
	"qlova.org/seed"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/css"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/sum"

	"qlova.org/seed/s/html/textarea"
)

//New returns a new textbox widget.
func New(text sum.String, options ...seed.Option) seed.Seed {
	_, variable := sum.ToString(text)

	var updater seed.Option

	switch v := variable.(type) {
	case *clientside.String:
		updater = seed.NewOption(func(c seed.Seed) {
			clientside.Hook(v, c)
			c.With(
				script.On("render", script.Element(c).Set("value", v)),
				script.On("input", v.SetTo(js.String{Value: script.Element(c).Get("value")})),
			)
		})
	case seed.Option:
		updater = v
	}

	return textarea.New(
		updater,
		css.SetResize(css.None),
		seed.Options(options),
	)
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
func Focus(c seed.Seed) script.Script {
	return func(q script.Ctx) {
		q(script.Element(c).Run(`focus`))
	}
}
