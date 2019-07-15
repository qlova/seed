package seed

import "github.com/qlova/seed/style/css"
import "github.com/qlova/seed/internal"

type Frame = internal.Frame
type Animation = internal.Animation

func (seed Seed) SetAnimation(animation Animation) {
	seed.animation = animation
	seed.SetAnimationName(css.AnimationName(seed.id))
	seed.SetAnimationDuration(css.Time(1))
	seed.SetAnimationIterationCount(css.Infinite)
}
