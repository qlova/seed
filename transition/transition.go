package transition

import (
	"fmt"
	"log"
	"reflect"

	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/page"
	"github.com/qlova/seed/popup"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/style/anime"
	"github.com/qlova/seed/view"
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
			log.Println("invalid seed type: ", reflect.TypeOf(c))
		}

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
