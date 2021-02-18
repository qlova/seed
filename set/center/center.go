package center

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/expander"
	"qlova.org/seed/set/visible"
)

//Seeds is a group of seeds to be centered.
type Seeds []seed.Seed

//AddTo implements seed.Option
func (s Seeds) AddTo(c seed.Seed) {
	expander.New().AddTo(c)
	for _, child := range s {
		child.AddTo(c)
	}
	expander.New().AddTo(c)
}

//This centers the provided seeds within their container..
func This(s ...seed.Seed) Seeds {
	return Seeds(s)
}

//ConditionalSeeds is a group of seeds to be centered upon a condition.
type ConditionalSeeds struct {
	Condition client.Bool
	Seeds
}

//AddTo implements seed.Option
func (s ConditionalSeeds) AddTo(c seed.Seed) {
	visible.When(s.Condition, expander.New()).AddTo(c)
	for _, child := range s.Seeds {
		child.AddTo(c)
	}
	visible.When(s.Condition, expander.New()).AddTo(c)
}

//When centers the given seeds within their container when the given condition is true.
func When(condition client.Bool, s ...seed.Seed) ConditionalSeeds {
	return ConditionalSeeds{condition, s}
}
