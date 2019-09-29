package style

import "github.com/qlova/seed/style/css"

//SetDelay sets the animation delay of the attached animation.
func (style Style) SetDelay(delay float64) {
	style.SetAnimationDelay(css.Time(delay))
}

//SetDuration sets the animation duration of the attached animation.
func (style Style) SetDuration(duration float64) {
	style.SetAnimationDuration(css.Time(duration))
}

//Reverse sets the attached animation to play in reverse.
func (style Style) Reverse() {
	style.SetAnimationDirection(css.Reverse)
}
