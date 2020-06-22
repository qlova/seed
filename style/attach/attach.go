package attach

import (
	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/style"
)

type Attacher interface {
	seed.Option
}

func ToParent() Attacher {
	return parentAttacher{}
}

type parentAttacher struct {
	style.Style
	rules css.Rules
}

func (parentAttacher) AddTo(c seed.Seed) {
	//var col = column.New()
	//d
}
