package feed_test

import (
	"image/color"

	"qlova.org/seed/web/app"
	"qlova.org/seed/client"
	"qlova.org/seed/web/js/console"
	"qlova.org/seed/new/button"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/text"
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
