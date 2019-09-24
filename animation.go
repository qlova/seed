package seed

import (
	"github.com/qlova/seed/internal"
	"github.com/qlova/seed/style/css"
)

//Frame is an animation frame.
type Frame = internal.Frame

//Animation is a change in styles across frames.
type Animation = internal.Animation

//SetAnimation sets the animation of this seed to be looping and 1 second long.
func (seed Seed) SetAnimation(animation Animation) {
	seed.animation = animation
	seed.SetAnimationName(css.AnimationName(seed.id))
	seed.SetAnimationDuration(css.Time(1))
	seed.SetAnimationIterationCount(css.Infinite)
}
