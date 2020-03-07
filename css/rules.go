package css

import (
	"github.com/qlova/seed"
)

type Rules []Rule

func (r Rules) Rules() Rules {
	return r
}

func (r Rules) AddTo(c seed.Any) {
	for _, rule := range r {
		rule.AddTo(c)
	}
}

func (r Rules) Apply(c seed.Ctx) {
	for _, rule := range r {
		rule.Apply(c)
	}
}

func (r Rules) Reset(c seed.Ctx) {
	for _, rule := range r {
		rule.Reset(c)
	}
}

func (r Rules) And(options ...seed.Option) seed.Option {
	return seed.And(r, options...)
}
