package script

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/html"
)

func init() {
	seed.Apply = func(s seed.Seed, ctx seed.Ctx) {
		s.Add(html.SetStyle("display", "none"))
		css.Set("display", "").Apply(s.Ctx(ctx.Ctx))
		ctx.Root().Add(s)
	}
	seed.Reset = func(s seed.Seed, ctx seed.Ctx) {
		css.Set("display", "none").Apply(s.Ctx(ctx.Ctx))
	}
}
