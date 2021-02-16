package hourbox

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/use/html"
	"qlova.org/seed/use/html/attr"
	"qlova.org/seed/use/js"

	"qlova.org/seed/new/textbox"
)

type duration struct {
	js.Value
}

func (d duration) GetDuration() js.Value {
	return d.Value
}

//New returns a new timebox widget.
func New(options ...seed.Option) seed.Seed {
	return textbox.New(attr.Set("type", "time"), seed.Options(options))
}

func init() {
	client.RegisterRootRenderer(func(seed.Seed) []byte {
		return []byte(`
		seeds.sethourbox = function(t) {
			if (t == 0) return null;
			let date = new Date(t);
			return date.getUTCHours().toString().padStart(2, 0) + ':' + 
				date.getUTCMinutes().toString().padStart(2, 0) + ':' + 
				date.getUTCSeconds().toString().padStart(2, 0);
		};
		seeds.gethourbox = function(t) {
			if (!t) return 0;

			let time = t.split(':');
			let duration = 0;
			duration += parseInt(time[0]) * 3.6e+6
			duration += parseInt(time[1]) * 60000
			if (time.length > 2) duration += parseInt(time[2]) * 1000;

			return duration;
		};
`)
	})
}

//Update updates the given variable whenever the hourbox time is modified.
//The duration returned is the number of milliseconds since the start of the day.
func Update(variable *clientside.Duration) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		clientside.Hook(variable, c)
		c.With(
			client.On("render", html.Element(c).Set("value", js.Func("seeds.sethourbox").Call(variable))),
			client.On("input", variable.SetTo(duration{js.Func("seeds.gethourbox").Call(html.Element(c).Get("value"))})),
		)
	})
}
