package css

import (
	"github.com/qlova/seed"
)

type Rules []Rule

func (r Rules) Rules() Rules {
	return r
}

func (r Rules) AddTo(c seed.Seed) {
	for _, rule := range r {
		rule.AddTo(c)
	}
}

func (r Rules) And(options ...seed.Option) seed.Option {
	return seed.And(r, options...)
}
