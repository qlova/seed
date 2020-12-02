package textbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"

	"qlova.org/seed/new/html/datalist"
	"qlova.org/seed/new/html/input"
	"qlova.org/seed/new/html/option"
	"qlova.org/seed/new/text"
)

//Data for a textbox.
type Data struct {
	Suggestions []string
}

//New returns a new textbox widget.
func New(options ...seed.Option) seed.Seed {
	Textbox := input.New(
		seed.Options(options),
	)

	var data Data
	Textbox.Load(&data)

	if len(data.Suggestions) > 0 {
		Datalist := datalist.New()

		for _, suggestion := range data.Suggestions {
			Datalist.With(
				option.New(text.SetString(suggestion)),
			)
		}

		Textbox.With(
			attr.Set("list", client.ID(Datalist)),
			Datalist,
		)
	}

	return Textbox
}

//Update updates the given variable whenever the textbox text is modified.
func Update(variable *clientside.String) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("value", variable)),
			client.On("input", variable.SetTo(js.String{Value: html.Element(c).Get("value")})),
		)
	})
}

//SetSuggestions sets suggestions for the value in the textbox.
func SetSuggestions(suggestions []string) seed.Option {
	return seed.Mutate(func(data *Data) {
		data.Suggestions = suggestions
	})
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
