package datebox

import (
	"time"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"

	"qlova.org/seed/new/textbox"
)

type date struct {
	js.Value
}

func (d date) GetTime() js.Value {
	return d.Value
}

//New returns a new datebox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "date"), seed.Options(options))
}

func init() {
	client.RegisterRootRenderer(func(seed.Seed) []byte {
		return []byte(`
		seeds.setdatebox = function(t) {
			if (t == -9223372036854) return null;
			let date = new Date(t);
			return date.getFullYear().toString() + '-' + (date.getMonth() + 1).toString().padStart(2, 0) + '-' + date.getDate().toString().padStart(2, 0);
		};
		seeds.getdatebox = function(t) {
			if (!t) return 0;

			let date = t.split('-');
			let time = new Date(parseInt(date[0]), parseInt(date[1])-1, parseInt(date[2]));

			return time.getTime();
		};
`)
	})
}

//Update updates the given variable whenever the datebox time is modified.
//The time is presented to the user in local time.
func Update(variable *clientside.Time) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("value", js.Func("seeds.setdatebox").Call(variable))),
			client.On("input", variable.SetTo(date{js.Func("seeds.getdatebox").Call(html.Element(c).Get("value"))})),
		)
	})
}

//SetMin sets minimum date range constraint
func SetMin(min time.Time) seed.Option {
	return attr.Set("min", min.Format("2006-01-02"))
}

//SetMinTo sets dynamic minimum date range constraint
func SetMinTo(min client.Time) seed.Option {
	return attr.SetTo("min", js.String{Value: js.Func("seeds.fdatebox").Call(min)})
}

//SetMax sets maximum date range constraint
func SetMax(max time.Time) seed.Option {
	return attr.Set("max", max.Format("2006-01-02"))
}

//SetMaxTo sets dynamic maximum date range constraint
func SetMaxTo(max client.Time) seed.Option {
	return attr.SetTo("min", js.String{Value: js.Func("seeds.fdatebox").Call(max)})
}
