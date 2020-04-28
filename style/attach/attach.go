package attach

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/s/column"
	"github.com/qlova/seed/style"
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
	var col = column.New()
	d
}
