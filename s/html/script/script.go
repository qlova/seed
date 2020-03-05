package script

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/html"
)

func New(options ...seed.Option) seed.Seed {
	return seed.New(html.SetTag("script").And(options...))
}
