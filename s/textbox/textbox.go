package textbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/sum"

	"qlova.org/seed/s/html/input"
)

//New returns a new textbox widget.
func New(text sum.String, options ...seed.Option) seed.Seed {
	_, variable := sum.ToString(text)

	var updater seed.Option

	switch v := variable.(type) {
	case *clientside.String:
		updater = Update(v)
	case seed.Option:
		updater = v
	}

	return input.New(
		updater,
		seed.Options(options),
	)
}

//Update updates the given variable whenever the textbox text is modified.
func Update(variable *clientside.String) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			script.On("render", script.Element(c).Set("value", variable)),
			script.On("input", variable.SetTo(js.String{Value: script.Element(c).Get("value")})),
		)
	})
}

//Var returns text with a variable text argument.
func Var(text state.String, options ...seed.Option) seed.Seed {
	if text.Null() {
		return New(nil, options...)
	}
	return New(seed.NewOption(func(c seed.Seed) {
		c.With(script.On("input", func(q script.Ctx) {
			text.Set(js.String{js.NewValue(script.Scope(c, q).Element() + `.value`)})(q)
		}), text.SetValue())
	}), seed.Options(options))
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
