package transition

import (
	"time"

	"qlova.org/seed/set"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/new/animation"
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
			0:   set.Translation(100%of.Parent, nil),
			100: set.Translation(nil, nil),
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
			0:   set.Translation(-100%of.Parent, nil),
			100: set.Translation(nil, nil),
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
			0:   set.Translation(nil, -100%of.Parent),
			100: set.Translation(nil, nil),
		},
		animation.Duration(400*time.Millisecond),
	)

	return New(
		In(slide),
		Out(slide.InReverse()),
	)
}
