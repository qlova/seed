package style

import "github.com/qlova/seed/style/css"

func (style Style) AnimateOffset(duration float64, justdelay ...float64) {
	var delay float64
	if len(justdelay) > 0 {
		delay = justdelay[0]
	}

	style.SetTransitionProperty((*css.Style).Top, (*css.Style).Left, (*css.Style).Right, (*css.Style).Bottom)
	style.SetWillChange((*css.Style).Top, (*css.Style).Left, (*css.Style).Right, (*css.Style).Bottom)

	style.SetTransitionDuration(css.Time(duration))
	style.SetTransitionDelay(css.Time(delay))
}

func (style Style) SetDelay(delay float64) {
	style.SetAnimationDelay(css.Time(delay))
} 

func (style Style) SetDuration(duration float64) {
	style.SetAnimationDuration(css.Time(duration))
}

func (style Style) Animate(duration float64, justdelay ...float64) {
	var delay float64
	if len(justdelay) > 0 {
		delay = justdelay[0]
	}

	style.SetTransitionProperty((*css.Style).Transform)
	style.SetWillChange((*css.Style).Transform)

	style.SetTransitionDuration(css.Time(duration))
	style.SetTransitionDelay(css.Time(delay))
}
