package text

import (
	"image/color"
	"strings"

	html_go "html"

	"qlova.org/seed"
	"qlova.org/seed/css"
	"qlova.org/seed/html"
	"qlova.org/seed/sum"
	"qlova.org/seed/unit"
)

//New returns a new text widget.
func New(text sum.String, options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("p"),
		html.SetInnerText(text),
		seed.Options(options),
	)
}

//Set sets the text content of the text.
func Set(value string) seed.Option {
	value = html_go.EscapeString(value)
	value = strings.Replace(value, "\n", "<br>", -1)
	value = strings.Replace(value, "  ", "&nbsp;&nbsp;", -1)
	value = strings.Replace(value, "\t", "&emsp;", -1)

	return html.Set(value)
}

//Color sets the color of the text.
func Color(c color.Color) css.Rule {
	return css.SetColor(css.RGB{Color: c})
}

//Size sets the font-size of the text.
func Size(s unit.Unit) css.Rule {
	return css.SetFontSize(css.Measure(s))
}

//Center aligns the text to to the center.
func Center() css.Rule {
	return css.SetTextAlign(css.Center)
}

//Right aligns the text to to the right.
func Right() css.Rule {
	return css.SetTextAlign(css.Right)
}

//Left aligns the text to to the left.
func Left() css.Rule {
	return css.SetTextAlign(css.Left)
}
