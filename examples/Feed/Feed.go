package main

import (
	"qlova.org/seed"
	"qlova.org/seed/app"
	"qlova.org/seed/js/console"
	"qlova.org/seed/s/button"
	"qlova.org/seed/s/feed"
	"qlova.org/seed/s/text"
	"qlova.org/seed/script"
	"qlova.org/seed/style/align"
	"qlova.org/seed/style/font"
)

func GetFeed() []string {
	return []string{"a", "b", "c"}
}

func main() {
	b := button.New("Click me")
	app.New("Feed",
		b,
		feed.New(GetFeed, feed.Do(func(f feed.Seed) {
			f.With(
				align.Center(),

				text.Var(f.Data.String()),
				text.New("hello",
					font.SetColor(seed.Red),
					script.OnClick(func(q script.Ctx) {
						console.Log(f.Data.String())
					}),
				),
			)
			b.With(script.OnClick(f.Refresh()))
		})),
	).Launch()
}
