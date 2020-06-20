package style

import (
	"qlova.org/seed/css"
)

type Attacher interface {
	Style

	Top() Attacher
	Left() Attacher
	Right() Attacher
	Bottom() Attacher
}

func AttachToParent() Attacher {
	return parentAttacher{}
}

type parentAttacher struct {
	Style
	rules css.Rules
}

func (p parentAttacher) Top() Attacher {
	p.rules = append(p.rules,
		css.SetTop(css.Zero),
		css.SetPosition(css.Absolute),
	)
	return parentAttacher{p.rules, p.rules}
}

func (p parentAttacher) Left() Attacher {
	p.rules = append(p.rules,
		css.SetLeft(css.Zero),
		css.SetPosition(css.Absolute),
	)
	return parentAttacher{p.rules, p.rules}
}

func (p parentAttacher) Right() Attacher {
	p.rules = append(p.rules,
		css.SetRight(css.Zero),
		css.SetPosition(css.Absolute),
	)
	return parentAttacher{p.rules, p.rules}
}

func (p parentAttacher) Bottom() Attacher {
	p.rules = append(p.rules,
		css.SetBottom(css.Zero),
		css.SetPosition(css.Absolute),
	)
	return parentAttacher{p.rules, p.rules}
}
