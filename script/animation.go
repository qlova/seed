package script

import (
	"github.com/qlova/seed/internal"
	"github.com/qlova/seed/style/css"
)

//Animation is a mapping of float64 to frames as a ratio between 0 and 1 (start and end).
type Animation = internal.Animation

//SetAnimation sets and plays the given animtion.
func (seed Seed) SetAnimation(animation *Animation) {
	if animation == nil {
		seed.Set("animation-name", "")
		return
	}
	var name = seed.Q.Context.Animation(animation)
	seed.Set("animation-name", name)
	seed.Set("animation-direction", "normal")
	seed.Set("animation-fill-mode", "forwards")
}

//SetAnimationReverse sets the animation to play in reverse.
func (seed Seed) SetAnimationReverse() {
	seed.Set("animation-direction", css.Reverse.String())
}

//SetAnimationDuration sets the duration that the animation should play for.
func (seed Seed) SetAnimationDuration(duration Float) {
	seed.SetDynamic("animation-duration", duration.LanguageType().Raw()+"+'s'")
}

//SetAnimationIterations sets how many times the animation should play.
//0 for infinite.
func (seed Seed) SetAnimationIterations(iterations Int) {
	seed.SetDynamic("animation-iteration-count", iterations.LanguageType().Raw())
}
