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

type Number float64

func (Number) numberValue() {}
func (n Number) Rule() Rule {
	return Rule(fmt.Sprintf("%f", n))
}

type Int int

func (Int) integerOrAutoValue()           {}
func (Int) animationIterationCountValue() {}
func (i Int) Rule() Rule {
	return Rule(fmt.Sprintf("%v", i))
}
