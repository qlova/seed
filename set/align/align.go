package align

import (
	"qlova.org/seed/web/css"
	"qlova.org/seed/set"
)

//Left aligns the seed to the start or left of its container.
func Left() css.Rule {
	return css.SetAlignSelf(css.FlexStart)
}

//Center center's the seed.
func Center() css.Rule {
	return css.SetAlignSelf(css.Center)
}

//Right aligns the seed to the end or right of its container.
func Right() css.Rule {
	return css.SetAlignSelf(css.FlexEnd)
}

//Bottom aligns the seed to the end or right of its container.
func Bottom() css.Rule {
	return css.SetAlignSelf(css.FlexEnd)
}

//Aligner can align things.
type Aligner interface {
	Left() set.Style
	Center() set.Style
	Right() set.Style
}

//Items returns an aligner that aligns children.
func Items() Aligner {
	return itemsAligner{}
}

type itemsAligner struct{}

func (itemsAligner) Left() set.Style {
	return css.SetAlignItems(css.FlexStart)
}

func (itemsAligner) Center() set.Style {
	return css.SetAlignItems(css.Center)
}

func (itemsAligner) Right() set.Style {
	return css.SetAlignItems(css.FlexStart)
}

//Text returns an aligner that aligns text.
func Text() Aligner {
	return textAligner{}
}

type textAligner struct{}

func (textAligner) Left() set.Style {
	return css.SetTextAlign(css.Left)
}

func (textAligner) Center() set.Style {
	return css.SetTextAlign(css.Center)
}

func (textAligner) Right() set.Style {
	return css.SetTextAlign(css.Right)
}
