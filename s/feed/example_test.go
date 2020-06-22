package feed_test

import (
	"image/color"

	"qlova.org/seed/app"
	"qlova.org/seed/client"
	"qlova.org/seed/js/console"
	"qlova.org/seed/s/button"
	"qlova.org/seed/s/feed"
	"qlova.org/seed/s/text"
	"qlova.org/seed/set/align"
)

func Example() {
	var values = feed.With(func() []string {
		return []string{"a", "b", "c"}
	})

	app.New("Feed",
		button.New("Click me", client.OnClick(values.Refresh())),

		values.New(align.Center(),
			text.New(values.Data),
			text.New("hello", text.Color(color.NRGBA{255, 0, 0, 255}),
				client.OnClick(console.Log(values.Data)),
			),
		),
	).Launch()
}
