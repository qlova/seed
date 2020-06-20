package main

import (
	"time"

	"qlova.org/seed"
	"qlova.org/seed/app"
	"qlova.org/seed/s/text"
	"qlova.org/seed/style"
	"qlova.org/seed/style/anime"
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
