package seed

import "github.com/qlova/seed/style"

type Animation style.Animation

func (seed Seed) SetOffsetAnimation(anim Animation) {
	seed.Style.SetOffsetAnimation(style.Animation(anim))
}
