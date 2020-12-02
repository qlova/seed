package spacing

import (
	"qlova.org/seed/use/css"
	"qlova.org/seed/set"
)

func Outside() set.Style {
	return css.SetJustifyContent(css.Center)
}

func Inbetween() set.Style {
	return css.SetJustifyContent(css.SpaceBetween)
}

func Before() set.Style {
	return css.SetJustifyContent(css.FlexEnd)
}

func After() set.Style {
	return css.SetJustifyContent(css.FlexStart)
}

func Divide() set.Style {
	return css.SetJustifyContent(css.SpaceEvenly)
}
