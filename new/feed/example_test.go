package feed_test

import (
	"image/color"

	"qlova.org/seed/client"
	"qlova.org/seed/new/app"
	"qlova.org/seed/new/button"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set/align"
	"qlova.org/seed/use/js"
	"qlova.org/seed/use/js/console"
)

func Example() {
	var values = feed.With(func() []string {
		return []string{"a", "b", "c"}
	})

	app.New("Feed",
		button.New(text.Set("Click me"), client.OnClick(values.Refresh())),

		values.New(align.Center(),
			text.New(text.SetStringTo(js.String{Value: values.Data.GetValue()})),
			text.New(text.Set("hello"), text.SetColor(color.NRGBA{255, 0, 0, 255}),
				client.OnClick(console.Log(values.Data)),
			),
		),
	).Launch()
}
