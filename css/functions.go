package css

import (
	"fmt"
	"time"
)

type Duration time.Duration

func (Duration) durationValue() {}
func (d Duration) Rule() Rule {
	return Rule(fmt.Sprintf("%fs", time.Duration(d).Seconds()))
}

type AnimationName string

func (AnimationName) animationNameValue() {}
func (a AnimationName) Rule() Rule {
	return Rule(a)
}
