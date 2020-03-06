package main

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/button"
	"github.com/qlova/seed/s/feed"
	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style/align"
)

func GetFeed() []string {
	return []string{"a", "b", "c"}
}

func main() {
	b := button.New("Click me")
	app.New("Feed",
		b,
		feed.New(GetFeed, feed.Do(func(f feed.Seed) {
			f.Add(
				align.Center(),

				text.Var(f.Data.String()),
				text.New("hello",
					text.SetColor(seed.Red),
					script.OnClick(func(q script.Ctx) {
						q.Print(f.Data.String())
					}),
				),

				//script.OnReady(f.Refresh()),

			)
			b.Add(script.OnClick(f.Refresh()))
		})),
	).Launch()
}
