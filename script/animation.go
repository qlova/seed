package script

import "github.com/qlova/seed/internal"
import "github.com/qlova/seed/style/css"

type Animation = internal.Animation

func (seed Seed) SetAnimation(animation *Animation) {
	var name = seed.Q.Context.Animation(animation)
	seed.Set("animation-name", name)
	seed.Set("animation-direction", "normal")
}

func (seed Seed) SetAnimationReverse() {
	seed.Set("animation-direction", css.Reverse.String())
}

func (seed Seed) SetAnimationDuration(duration Float) {
	seed.SetDynamic("animation-duration", duration.LanguageType().Raw()+"+'s'")
}

func (seed Seed) SetAnimationIterations(iterations Int) {
	seed.SetDynamic("animation-iteration-count", iterations.LanguageType().Raw())
}
