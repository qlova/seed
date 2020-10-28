package slider

import (
	"strconv"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"

	"qlova.org/seed/new/html/input"
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
			client.On("render", html.Element(c).Set("value", variable)),
			client.On("input", variable.SetTo(js.Number{Value: js.NewValue("+%v", html.Element(c).Get("value"))})),
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
func Focus(c seed.Seed) js.Script {
	return func(q js.Ctx) {
		q(html.Element(c).Run(`focus`))
	}
}
