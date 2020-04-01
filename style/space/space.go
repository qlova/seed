package space

import (
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/style"
)

//Spacer can align things.
type Spacer interface {
	Outside() style.Style
}

//Items returns an aligner that aligns children.
func Items() Spacer {
	return itemsSpacer{}
}

type itemsSpacer struct{}

func (itemsSpacer) Outside() style.Style {
	return css.SetJustifyContent(css.Center)
}
