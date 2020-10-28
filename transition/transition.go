package transition

import (
	"fmt"
	"time"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/css"
	"qlova.org/seed/js"
	"qlova.org/seed/vfx/animation"
)

var fadeIn = animation.New(
	animation.Frames{
		0:   css.SetOpacity(css.Zero),
		100: css.SetOpacity(css.Number(1)),
	},
	animation.Duration(400*time.Millisecond),
)

var fadeOut = animation.New(
	animation.Frames{
		0:   css.SetOpacity(css.Number(1)),
		100: css.SetOpacity(css.Zero),
	},
	animation.Duration(400*time.Millisecond),
)

type Transition struct {
	seed.Option

	In, Out animation.Animation
}

type data struct {
	OnEnter, OnExit func(...client.Script) seed.Option
}

type Option func(*Transition)

func SetOnEnter(onenter func(...client.Script) seed.Option) seed.Option {
	return seed.Mutate(func(d *data) {
		d.OnEnter = onenter
	})
}

func SetOnExit(onexit func(...client.Script) seed.Option) seed.Option {
	return seed.Mutate(func(d *data) {
		d.OnExit = onexit
	})
}

func New(options ...Option) Transition {
	var t Transition
	for _, o := range options {
		o(&t)
	}

	t.Option = seed.NewOption(func(c seed.Seed) {

		enter := js.Script(func(q js.Ctx) {
			client.Option(t.In, q).AddTo(c)
			fmt.Fprintf(q, `seed.in(%v, 0.4);`, client.Element(c))
		})

		exit := js.Script(func(q js.Ctx) {
			client.Option(t.Out, q).AddTo(c)
			fmt.Fprintf(q, `seed.out(%v, 0.4);`, client.Element(c))
		})

		var d data
		c.Load(&d)

		if d.OnEnter != nil && d.OnExit != nil {
			c.With(
				d.OnEnter(enter),
				d.OnExit(exit),
			)
		} else {
			c.With(
				client.On("visible", js.Script(func(q js.Ctx) {
					client.Option(t.In, q).AddTo(c)

				})),
				client.On("hidden", js.Script(func(q js.Ctx) {
					client.Option(t.Out, q).AddTo(c)
				})),
			)
		}
	})

	return t
}

func In(in animation.Animation) Option {
	return func(t *Transition) {
		t.In = in
	}
}

func Out(out animation.Animation) Option {
	return func(t *Transition) {
		t.Out = out
	}
}
