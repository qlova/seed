package numberbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/s/html/input"
	"qlova.org/seed/script"
	"qlova.org/seed/state"
	"qlova.org/seed/sum"
)

//New returns a new textbox widget.
func New(s sum.Float64, options ...seed.Option) seed.Seed {
	_, variable := sum.ToFloat64(s)

	var updater seed.Option

	switch v := variable.(type) {
	case *clientside.Float64:
		updater = seed.NewOption(func(c seed.Seed) {
			clientside.Hook(v, c)
			c.With(
				script.On("render", script.Element(c).Set("value", v)),
				script.On("input", v.SetTo(js.Number{Value: js.Func("Number").Call(script.Element(c).Get("value"))})),
			)
		})
	case seed.Option:
		updater = v
	}

	return input.New(
		attr.Set("type", "number"),
		updater,
		seed.Options(options),
	)
}

//Var returns text with a variable text argument.
func Var(f state.Float, options ...seed.Option) seed.Seed {
	if f.Null() {
		return New(nil, options...)
	}
	return New(seed.NewOption(func(c seed.Seed) {
		c.With(script.On("input", func(q script.Ctx) {
			f.Set(js.Number{js.NewValue(script.Scope(c, q).Element() + `.value`)})(q)
		}), state.SetProperty("value", f))
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
