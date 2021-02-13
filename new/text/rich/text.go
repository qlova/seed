package rich

import (
	"html"
	"image/color"
	"strings"
)

const (
	rich = iota

	style

	reset

	italic
	bold
	underline
	strikethrough

	rgb
	rgba

	font
	icon
	link
)

//Text holds rich-formatted text.
type Text string

//Plain returns true if the text has no formatting applied to it.
func (text Text) Plain() bool {
	for _, char := range text {
		if char == rich || char == style {
			return false
		}
	}
	return true
}

func (text *Text) format() {
	if len(*text) > 0 && (*text)[0] == rich {
		*text = (*text)[1:]
	}
}

//Italic returns the text formatted as italic.
func (text Text) Italic() Text {
	text.format()
	return Text([]byte{rich, style, italic}) + text + Text([]byte{rich, style, reset})
}

//Bold returns the text formatted bold.
func (text Text) Bold() Text {
	text.format()
	return Text([]byte{rich, style, bold}) + text + Text([]byte{rich, style, reset})
}

//In returns the text in the given color.
func (text Text) In(c color.Color) Text {
	text.format()

	r, g, b, a := c.RGBA()
	if a == 0xffff {

		return Text([]byte{rich, style, rgb}) + Text(encode3ToString(uint8(r>>8), uint8(g>>8), uint8(b>>8))) + text + Text([]byte{rich, style, reset})
	}
	if a == 0 {
		return Text([]byte{rich, style, rgba}) + Text(encode4ToString(0, 0, 0, 0)) + text + Text([]byte{rich, style, reset})
	}

	// Since Color.RGBA returns an alpha-premultiplied color, we should have r <= a && g <= a && b <= a.
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a

	return Text([]byte{rich, style, rgba}) + Text(encode4ToString(uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))) + text + Text([]byte{rich, style, reset})
}

//Icon embeds an icon in the text.
func Icon(src string) Text {
	if len(src) > 255 {
		panic("rich.Icon src length less than 255, use a shorter length or fix the rich package.")
	}
	return Text(append([]byte{rich, style, icon, byte(len(src))}, src...)) + Text([]byte{rich, style, reset})
}

//Link embeds a url link in the text with the given label.
//If the label is the empty string, the url is used as the link.
func Link(url, label string) Text {
	if len(url) > 255 {
		panic("rich.Link src length less than 255, use a shorter length or fix the rich package.")
	}
	return Text(append([]byte{rich, style, link, byte(len(url)), byte(len(label))}, (url+string(label))...)) + Text([]byte{rich, style, reset})
}

func (text Text) String() string {
	if text.Plain() {
		return string(text)
	}

	parse := func(s string) string {
		switch text[0] {
		case rgb:
			return text[7:].String()
		case rgba:
			return text[9:].String()
		case font:
			return text[text[1]+2:].String()
		case icon:
			return text[text[1]+2:].String()
		case link:
			return text[3+text[1]:].String()
		}

		return text[2:].String()
	}

	var splits = strings.Split(string(text), string([]byte{rich}))
	var result string

	for _, split := range splits {
		result += parse(split)
	}

	return result
}

//HTML returns the text formatted as HTML.
func (text Text) HTML() string {

	var convert func(s string) string
	convert = func(s string) string {
		if len(s) == 0 {
			return ""
		}

		if s[0] > 1 {
			s := html.EscapeString(string(s))
			s = strings.Replace(s, "\n", "<br>", -1)
			s = strings.Replace(s, "  ", "&nbsp;&nbsp;", -1)
			s = strings.Replace(s, "\t", "&emsp;", -1)

			if s[len(s)-1] == ' ' {
				s = s[:len(s)-1] + "&nbsp;"
			}

			return s
		}

		switch s[1] {
		case reset:
			return s[2:]
		case italic:
			return "<em>" + convert(s[2:]) + "</em>"
		case bold:
			return "<strong>" + convert(s[2:]) + "</strong>"
		case rgb:
			return "<span style='color:#" + s[2:8] + ";'>" + convert(s[8:]) + "</span>"
		case rgba:
			return "<span style='color:#" + s[2:10] + ";'>" + convert(s[10:]) + "</span>"
		case icon:
			return "<img style='margin-top: 0.1em;vertical-align:text-top;height:1em;font-size:inherit;' src='" + s[3:3+int(s[2])] + "'>" + convert(s[3+int(s[2]):])
		case link:
			url := s[4 : 3+int(s[2])]
			label := convert(s[4+int(s[2]) : 4+int(s[2])+int(s[3])])
			return "<a href='" + url + "'>" + label + "</a>"
		default:
			panic("invalid text format")
		}
	}

	if text.Plain() {
		return convert(string(text))
	}

	var splits = strings.Split(string(text), string([]byte{rich}))
	var result string

	if len(splits) == 0 {
		return convert(string(text))
	}

	for _, split := range splits {
		result += convert(split)
	}

	return result
}
