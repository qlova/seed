package transition

import (
	"fmt"
	"time"

	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/page"
	"qlova.org/seed/popup"
	"qlova.org/seed/script"
	"qlova.org/seed/vfx/animation"
	"qlova.org/seed/view"
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

type Option func(*Transition)

func New(options ...Option) Transition {
	var t Transition
	for _, o := range options {
		o(&t)
	}

	t.Option = seed.NewOption(func(c seed.Seed) {

		enter := func(q script.Ctx) {
			t.In.AddTo(script.Scope(c, q))
			fmt.Fprintf(q, `seed.in(%v, 0.4);`, script.Scope(c, q).Element())
		}

		exit := func(q script.Ctx) {
			t.Out.AddTo(script.Scope(c, q))
			fmt.Fprintf(q, `seed.out(%v, 0.4);`, script.Scope(c, q).Element())
		}

		switch c.(type) {
		case page.Seed:
			c.With(
				page.OnEnter(enter),
				page.OnExit(exit),
			)
		case popup.Seed:
			c.With(
				popup.OnShow(enter),
				popup.OnHide(exit),
			)
		case view.Seed:
			c.With(
				view.OnEnter(enter),
				view.OnExit(exit),
			)
		default:
			c.With(
				script.On("visible", func(q script.Ctx) {
					t.In.AddTo(script.Scope(c, q))

				}),
				script.On("hidden", func(q script.Ctx) {
					t.Out.AddTo(script.Scope(c, q))
				}),
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
