package transition

import (
	"github.com/qlova/seed/style"
	"github.com/qlova/seed/style/anime"
)

//SlideRight slides in from the right and then out to the left.
func SlideRight() Transition {

	slide := anime.New(
		anime.Keyframes{
			0:   style.Translate(100, 0),
			100: style.Translate(0, 0),
		},
	)

	return New(
		In(slide),
		Out(slide.Reverse()),
	)
}
