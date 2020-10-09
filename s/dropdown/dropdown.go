//Package dropdown provides an input widget for selecting from an array of preset values.
package dropdown

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/html"
	"qlova.org/seed/html/attr"
	"qlova.org/seed/js"
	"qlova.org/seed/s/html/htmlselect"
	"qlova.org/seed/s/html/option"
	"qlova.org/seed/s/text"
	"qlova.org/seed/s/text/rich"
)

type data struct {
	seed.Data

	values []string

	placeholder string
}

//New returns a new dropdown.
func New(options ...seed.Option) seed.Seed {

	var data data

	var c = htmlselect.New(options...)

	c.Read(&data)

	if data.placeholder != "" {
		option.New(
			text.Set(rich.Text(data.placeholder)),
			attr.Set("disabled", ""),
			attr.Set("selected", ""),
			attr.Set("hidden", ""),
			attr.Set("value", ""),
		).AddTo(c)
	}

	for _, val := range data.values {
		option.New(text.Set(rich.Text(val))).AddTo(c)
	}

	return c
}

//Update updates the given variable whenever the dropdown value is modified.
func Update(variable *clientside.String) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("value", variable)),
			client.On("input", variable.SetTo(js.String{Value: html.Element(c).Get("value")})),
		)
	})
}

//Set sets the preset dropdown values.
func Set(values []string) seed.Option {
	return seed.Mutate(func(d *data) {
		d.values = values
	})
}

//SetPlaceholder sets the placeholder text of the dropbox.
func SetPlaceholder(placeholder string) seed.Option {
	return seed.Mutate(func(d *data) {
		d.placeholder = placeholder
	})
}
