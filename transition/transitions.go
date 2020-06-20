package transition

import (
	"qlova.org/seed/style"
	"qlova.org/seed/style/anime"
)

func Fade() Transition {
	return New(
		In(fadeIn),
		Out(fadeOut),
	)
}

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

//SlideLeft slides in from the left and then out to the right.
func SlideLeft() Transition {

	slide := anime.New(
		anime.Keyframes{
			0:   style.Translate(-100, 0),
			100: style.Translate(0, 0),
		},
	)

	return New(
		In(slide),
		Out(slide.Reverse()),
	)
}

//DropDown slides in from the right and then out to the left.
func DropDown() Transition {

	slide := anime.New(
		anime.Keyframes{
			0:   style.Translate(0, -100),
			100: style.Translate(0, 0),
		},
	)

	return New(
		In(slide),
		Out(slide.Reverse()),
	)
}
