package space

import (
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/style"
)

//Spacer can align things.
type Spacer interface {
	Outside() style.Style
	Inside() style.Style
	Before() style.Style
	Divide() style.Style
}

//Items returns an aligner that aligns children.
func Items() Spacer {
	return itemsSpacer{}
}

type itemsSpacer struct{}

func (itemsSpacer) Outside() style.Style {
	return css.SetJustifyContent(css.Center)
}

func (itemsSpacer) Inside() style.Style {
	return css.SetJustifyContent(css.SpaceBetween)
}

func (itemsSpacer) Before() style.Style {
	return css.SetJustifyContent(css.FlexEnd)
}

func (itemsSpacer) Divide() style.Style {
	return css.SetJustifyContent(css.SpaceEvenly)
}
