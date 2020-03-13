package align

import (
	"github.com/qlova/seed/css"
	"github.com/qlova/seed/style"
)

//Center center's the seed.
func Center() css.Rule {
	return css.SetAlignSelf(css.Center)
}

//Aligner can align things.
type Aligner interface {
	Left() style.Style
	Center() style.Style
	Right() style.Style
}

//Items returns an aligner that aligns children.
func Items() Aligner {
	return itemsAligner{}
}

type itemsAligner struct{}

func (itemsAligner) Left() style.Style {
	return css.SetAlignItems(css.FlexStart)
}

func (itemsAligner) Center() style.Style {
	return css.SetAlignItems(css.Center)
}

func (itemsAligner) Right() style.Style {
	return css.SetAlignItems(css.FlexStart)
}

//Text returns an aligner that aligns text.
func Text() Aligner {
	return textAligner{}
}

type textAligner struct{}

func (textAligner) Left() style.Style {
	return css.SetTextAlign(css.Left)
}

func (textAligner) Center() style.Style {
	return css.SetTextAlign(css.Center)
}

func (textAligner) Right() style.Style {
	return css.SetTextAlign(css.Right)
}
