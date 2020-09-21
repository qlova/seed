package spacing

import (
	"qlova.org/seed/css"
	"qlova.org/seed/style"
)

func Outside() style.Style {
	return css.SetJustifyContent(css.Center)
}

func Inbetween() style.Style {
	return css.SetJustifyContent(css.SpaceBetween)
}

func Before() style.Style {
	return css.SetJustifyContent(css.FlexEnd)
}

func After() style.Style {
	return css.SetJustifyContent(css.FlexStart)
}

func Divide() style.Style {
	return css.SetJustifyContent(css.SpaceEvenly)
}
