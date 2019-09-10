package css

import (
	"bytes"
)

type unicodeRange string

func (u unicodeRange) String() string { return string(u) }

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
func (font FontFace) String() string   { return font.FontFamily }

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

func (font FontFace) Get(property string) string {
	switch property {
	case "font-family":
		return font.FontFamily
	case "src":
		return font.Src.String()
	case "font-stretch":
		return font.FontStretch.String()
	case "font-style":
		return font.FontStyle.String()
	case "font-weight":
		return font.FontWeight.String()
	case "unicode-range":
		return font.UnicodeRange.String()
	case "font-display":
		return font.FontDisplay.String()
	}
	return ""
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
		buffer.WriteString(font.FontStretch.String())
		buffer.WriteByte(';')
	}

	if font.FontStyle != nil {
		buffer.WriteString("font-style: ")
		buffer.WriteString(font.FontStyle.String())
		buffer.WriteByte(';')
	}

	if font.FontWeight != nil {
		buffer.WriteString("font-weight: ")
		buffer.WriteString(font.FontWeight.String())
		buffer.WriteByte(';')
	}

	if font.UnicodeRange.String() != "" {
		buffer.WriteString("unicode-range: ")
		buffer.WriteString(font.UnicodeRange.String())
		buffer.WriteByte(';')
	}

	if font.FontDisplay.String() != "" {
		buffer.WriteString("font-display: ")
		buffer.WriteString(font.FontDisplay.String())
		buffer.WriteByte(';')
	}

	return buffer.Bytes()
}
