package center

import (
	"qlova.org/seed"
	"qlova.org/seed/s/expander"
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
