package css

import (
	"bytes"
)

type unicodeRange string

func (u unicodeRange) Rule() Rule { return Rule(u) }

type FontFace struct {
	FontFamily   string
	Src          urlType
	FontStretch  fontStretchValue
	FontStyle    fontStyleValue
	FontWeight   fontWeightValue
	FontDisplay  fontDisplayValue
	UnicodeRange unicodeRange
}

func (font FontFace) fontFamilyValue() {}
func (font FontFace) Rule() Rule       { return Rule(font.FontFamily) }

func NewFontFace(name string, src string) FontFace {
	return FontFace{
		FontFamily:  name,
		Src:         urlType("url(" + src + ")"),
		FontDisplay: Swap,
	}
}

func (font FontFace) Set(property, value string) {
	switch property {
	case "font-family":
		font.FontFamily = value
	case "src":
		font.Src = urlType(value)
	case "font-stretch":
		font.FontStretch = fontStretchType(value)
	case "font-style":
		font.FontStyle = fontStyleType(value)
	case "font-weight":
		font.FontWeight = fontWeightType(value)
	case "unicode-range":
		font.UnicodeRange = unicodeRange(value)
	case "font-display":
		font.FontDisplay = fontDisplayType(value)
	}
}

func (font FontFace) Bytes() []byte {
	var buffer bytes.Buffer
	buffer.WriteString("font-family:")
	buffer.WriteString(font.FontFamily)
	buffer.WriteByte(';')

	buffer.WriteString("src:")
	buffer.WriteString(font.Src.String())
	buffer.WriteByte(';')

	if font.FontStretch != nil {
		buffer.WriteString("font-stretch: ")
		buffer.WriteString(string(font.FontStretch.Rule()))
		buffer.WriteByte(';')
	}

	if font.FontStyle != nil {
		buffer.WriteString("font-style: ")
		buffer.WriteString(string(font.FontStyle.Rule()))
		buffer.WriteByte(';')
	}

	if font.FontWeight != nil {
		buffer.WriteString("font-weight: ")
		buffer.WriteString(string(font.FontWeight.Rule()))
		buffer.WriteByte(';')
	}

	if font.UnicodeRange.Rule() != "" {
		buffer.WriteString("unicode-range: ")
		buffer.WriteString(string(font.UnicodeRange.Rule()))
		buffer.WriteByte(';')
	}

	if font.FontDisplay.Rule() != "" {
		buffer.WriteString("font-display: ")
		buffer.WriteString(string(font.FontDisplay.Rule()))
		buffer.WriteByte(';')
	}

	return buffer.Bytes()
}
