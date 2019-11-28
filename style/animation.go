package style

import "github.com/qlova/seed/style/css"

//Animate animates this elements transform.
func (style Style) Animate(duration float64) {
	style.CSS().Set("transition-property", "transform")
	style.CSS().SetWillChange((*css.Style).Transform)
	style.CSS().SetTransitionDuration(css.Time(duration))
}

//SetDelay sets the animation delay of the attached animation.
func (style Style) SetDelay(delay float64) {
	style.CSS().SetAnimationDelay(css.Time(delay))
}

//SetDuration sets the animation duration of the attached animation.
func (style Style) SetDuration(duration float64) {
	style.CSS().SetAnimationDuration(css.Time(duration))
}

//Reverse sets the attached animation to play in reverse.
func (style Style) Reverse() {
	style.CSS().SetAnimationDirection(css.Reverse)
}
