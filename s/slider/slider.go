package slider

import (
	"strconv"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/script"

	"qlova.org/seed/s/html/input"
)

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	return input.New(attr.Set("type", "range"), seed.Options(options))
}

//Update updates the given variable whenever the textbox text is modified.
func Update(variable *clientside.Int) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", script.Element(c).Set("value", variable)),
			client.On("input", variable.SetTo(js.Number{Value: script.Element(c).Get("value")})),
		)
	})
}

//SetMin sets the minumum value of this slider.
func SetMin(min int) seed.Option {
	return attr.Set("min", strconv.Itoa(min))
}

//SetMax sets the maximum value of this slider.
func SetMax(max int) seed.Option {
	return attr.Set("max", strconv.Itoa(max))
}

//SetRange sets the range of the slider.
func SetRange(min, max int) seed.Option {
	return seed.Options{
		SetMin(min), SetMax(max),
	}
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
