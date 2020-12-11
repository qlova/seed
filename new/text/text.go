package text

import (
	"fmt"
	"image/color"
	"strings"

	html_go "html"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/text/rich"
	"qlova.org/seed/use/css"
	"qlova.org/seed/use/css/units"
	"qlova.org/seed/use/html"
)

//New returns a new text widget.
func New(options ...seed.Option) seed.Seed {
	return seed.New(
		html.SetTag("span"),
		seed.Options(options),
	)
}

//SetString sets the text from the given string.
func SetString(value string) seed.Option {
	value = html_go.EscapeString(value)
	value = strings.Replace(value, "\n", "<br>", -1)
	value = strings.Replace(value, "  ", "&nbsp;&nbsp;", -1)
	value = strings.Replace(value, "\t", "&emsp;", -1)

	return html.Set(value)
}

//Set sets the text content of the text to the given formatted lines, each argument is seperated by a newline.
func Set(lines ...rich.Text) seed.Option {

	var result string

	for i, line := range lines {
		if i > 0 {
			result += "<br>"
		}
		result += line.HTML()
	}

	return html.Set(result)
}

//SetStringTo sets the text content of the text.
func SetStringTo(value client.String) seed.Option {
	return html.SetInnerTextTo(value)
}

//SetColor sets the color of the text.
func SetColor(c color.Color) css.Rule {
	return css.SetColor(css.RGB{Color: c})
}

//SetSize sets the font-size of the text.
func SetSize(s units.Unit) css.Rule {
	return css.SetFontSize(css.Measure(s))
}

//SetLineHeight sets the line-height of the text.
func SetLineHeight(height float64) css.Rule {
	return css.Set("line-height", fmt.Sprint(height))
}

//Selectable sets the text content to be selectable.
func Selectable() seed.Option {
	return css.SetUserSelect(css.Text)
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
