package transition

import (
	"time"

	"qlova.org/seed/style"
	"qlova.org/seed/vfx/animation"
)

func Fade() Transition {
	return New(
		In(fadeIn),
		Out(fadeOut),
	)
}

//SlideRight slides in from the right and then out to the left.
func SlideRight() Transition {

	slide := animation.New(
		animation.Frames{
			0:   style.Translate(100, 0),
			100: style.Translate(0, 0),
		},
		animation.Duration(400*time.Millisecond),
	)

	return New(
		In(slide),
		Out(slide.InReverse()),
	)
}

//SlideLeft slides in from the left and then out to the right.
func SlideLeft() Transition {

	slide := animation.New(
		animation.Frames{
			0:   style.Translate(-100, 0),
			100: style.Translate(0, 0),
		},
		animation.Duration(400*time.Millisecond),
	)

	return New(
		In(slide),
		Out(slide.InReverse()),
	)
}

//DropDown slides in from the right and then out to the left.
func DropDown() Transition {

	slide := animation.New(
		animation.Frames{
			0:   style.Translate(0, -100),
			100: style.Translate(0, 0),
		},
		animation.Duration(400*time.Millisecond),
	)

	return New(
		In(slide),
		Out(slide.InReverse()),
	)
}
