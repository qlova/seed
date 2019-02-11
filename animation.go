package seed

import "github.com/qlova/seed/style"
import "github.com/qlova/seed/style/css"

type Frame struct {
	style.Style
}
type Animation map[float64]func(Frame)

func (seed Seed) SetAnimation(animation Animation) {
	seed.animation = animation
	seed.SetAnimationName(css.AnimationName(seed.id))
	seed.SetAnimationDuration(css.Time(1))
	seed.SetAnimationIterationCount(css.Infinite)
}

func (animation Animation) Bytes() []byte {
	var Keyframes = make(css.Keyframes, len(animation))
	for time := range animation {
		var frame = Frame{Style: style.New()}
		animation[time](frame)

		println(frame.Style.Style.Left().String())
		
		Keyframes[time] = frame.Style.Style
	}
	return Keyframes.Bytes()
}
