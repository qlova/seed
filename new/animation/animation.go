package animation

import (
	"strconv"
	"time"

	"qlova.org/seed"
	"qlova.org/seed/use/css"
)

var id int

//Animation can be added to a seed as the current playing animation of that seed.
type Animation struct {
	id int

	frames  Frames
	options []seed.Option
}

//InReverse returns a new animation which is the reverse of anim.
func (anim Animation) InReverse() Animation {
	Reverse().addto(&anim)
	return anim
}

func Set(anim Animation) seed.Option {
	return seed.NewOption(func(c seed.Seed) {
		var data data
		c.Load(&data)

		data.animations = append(data.animations, anim)

		c.Save(data)

		css.SetAnimationName(css.AnimationName("a" + strconv.Itoa(anim.id))).AddTo(c)
		for _, o := range anim.options {
			o.AddTo(c)
		}
	})
}

type data struct {
	

	animations []Animation
}

//Frame is a individual keyframe of an animation.
type Frame []css.Style

//Rules implements css.Style
func (f Frame) Rules() css.Rules {
	var result css.Rules
	for _, s := range f {
		result = append(result, s.Rules()...)
	}
	return result
}

//Frames is a mapping between time and a keyframe.
type Frames map[float64]css.Style

func (f Frames) addto(a *Animation) {
	a.frames = f
}

//Option values can be passed to New to configure the animation.
type Option interface {
	addto(a *Animation)
}

//NewOption can be used to create new Options.
type NewOption func(*Animation)

func (o NewOption) addto(a *Animation) { o(a) }

//New returns a new animation.
//It also takes a variable amount of options that will be applied to any seed with this animation set.
//These are useful for setting default animation durations, looping etc.
func New(options ...Option) Animation {
	id++

	var a = Animation{id: id}

	for _, option := range options {
		option.addto(&a)
	}

	return a
}

//Loop sets the current playing animation to loop back and forth.
func Loop() Option {
	return NewOption(func(a *Animation) {
		a.options = append(a.options, css.Rules{
			css.SetAnimationDirection(css.Alternate),
			css.SetAnimationIterationCount(css.Infinite),
		})
	})
}

//Reverse sets the current playing animation to play in reverse.
func Reverse() Option {
	return NewOption(func(a *Animation) {
		a.options = append(a.options, css.Rules{
			css.SetAnimationDirection(css.AlternateReverse),
		})
	})
}

//Duration sets the duration of the current playing animation.
func Duration(d time.Duration) Option {
	return NewOption(func(a *Animation) {
		a.options = append(a.options, css.Rules{
			css.SetAnimationDuration(css.Duration(d)),
		})
	})
}

func (anim Animation) AddTo(c seed.Seed) {
	var data data
	c.Load(&data)

	data.animations = append(data.animations, anim)

	c.Save(data)

	css.SetAnimationName(css.AnimationName("a" + strconv.Itoa(anim.id))).AddTo(c)
	for _, o := range anim.options {
		o.AddTo(c)
	}
}
