package style

import "github.com/qlova/seed/style/css"

type Animation struct {
	Duration, Delay float64
}

func (style *Style) SetOffsetAnimation(animation Animation) {
	style.SetTransitionProperty((*css.Style).Left, (*css.Style).Top, (*css.Style).Right, (*css.Style).Bottom)
	if animation.Duration > 0 {
		style.SetTransitionDuration(css.Time(animation.Duration))
	}
	if animation.Delay > 0 {
		style.SetTransitionDuration(css.Time(animation.Delay))
	}
}
