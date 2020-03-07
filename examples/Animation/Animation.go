package main

import (
	"time"

	"github.com/qlova/seed"
	"github.com/qlova/seed/app"
	"github.com/qlova/seed/s/text"
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/anime"
)

func main() {

	var MoveRightAndLeft = anime.New(
		anime.Keyframes{
			0: anime.Set(
				style.SetTextColor(seed.Red),
				style.Translate(0, 0),
			),
			100: anime.Set(
				style.Translate(100, 0),
			),
		},
		anime.Loop(),
		anime.SetDuration(2*time.Second),
	)

	app.New("Animation",
		text.New("Hello", MoveRightAndLeft),
	).Launch()
}
