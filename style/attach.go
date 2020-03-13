package style

import (
	"github.com/qlova/seed/css"
)

type Attacher interface {
	Top() Style
	Left() Style
	Right() Style
	Bottom() Style
}

func AttachToParent() Attacher {
	return parentAttacher{}
}

type parentAttacher struct{}

func (parentAttacher) Top() Style {
	return css.Rules{
		css.SetTop(css.Zero),
		css.SetPosition(css.Absolute),
	}
}

func (parentAttacher) Left() Style {
	return css.Rules{
		css.SetLeft(css.Zero),
		css.SetPosition(css.Absolute),
	}
}

func (parentAttacher) Right() Style {
	return css.Rules{
		css.SetRight(css.Zero),
		css.SetPosition(css.Absolute),
	}
}

func (parentAttacher) Bottom() Style {
	return css.Rules{
		css.SetBottom(css.Zero),
		css.SetPosition(css.Absolute),
	}
}
