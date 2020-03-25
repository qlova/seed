package transition

import (
	"fmt"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style/anime"
)

var fadeIn = anime.New(
	anime.Keyframes{
		0:   css.SetOpacity(css.Zero),
		100: css.SetOpacity(css.Number(1)),
	},
)

var fadeOut = anime.New(
	anime.Keyframes{
		0:   css.SetOpacity(css.Number(1)),
		100: css.SetOpacity(css.Zero),
	},
)

type Transition struct {
	seed.Option

	In, Out anime.Animation
}

type Option func(*Transition)

func New(options ...Option) Transition {
	var t Transition
	for _, o := range options {
		o(&t)
	}

	t.Option = seed.Do(func(c seed.Seed) {
		c.Add(
			page.OnEnter(func(q script.Ctx) {
				t.In.AddTo(q.Scope(c))
				fmt.Fprintf(q, `seed.in(%v, 0.4);`, q.Scope(c).Element())
			}).And(page.OnExit(func(q script.Ctx) {
				t.Out.AddTo(q.Scope(c))
				fmt.Fprintf(q, `seed.out(%v, 0.4);`, q.Scope(c).Element())
			})),
		)
	})

	return t
}

func In(in anime.Animation) Option {
	return func(t *Transition) {
		t.In = in
	}
}

func Out(out anime.Animation) Option {
	return func(t *Transition) {
		t.Out = out
	}
}

func Fade() Transition {
	return New(
		In(fadeIn),
		Out(fadeOut),
	)
}
