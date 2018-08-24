package app

import "bytes"

type Style struct {
	css bytes.Buffer
}

func (style *Style) Set(property, value string) {
	style.css.WriteString(property)
	style.css.WriteByte(':')
	style.css.WriteString(value)
	style.css.WriteByte(';')
}

func (style *Style) SetBackgroundColor(color string) {
	style.Set("background-color", color)
}

func (style *Style) SetLayout(layout string) {
	style.Set("display", layout)
}

func (style *Style) AutoExpand() {
	style.Set("flex-grow", "1")
}

func (style *Style) SetPosition(x, y string) {
	 style.Set("top", y)
	 style.Set("left", x)
}

func (style *Style) SetSize(width, height string) {
	 style.Set("width", width)
	 style.Set("height", height)
}

func (style *Style) SetMaxHeight(height string) {
	 style.Set("max-height", height)
}

func (style *Style) SetHeight(height string) {
	 style.Set("height", height)
}

func (style *Style) SetWidth(width string) {
	 style.Set("width", width)
}

func (style *Style) SetSticky() {
	 style.Set("position", "fixed")
}

func (style *Style) SetHidden() {
	 style.Set("display", "none")
}

func (style *Style) SetVisible() {
	 style.Set("display", "initial")
}

func (style *Style) AttachLeft() {
	 style.Set("left", "0")
}

func (style *Style) AttachBottom() {
	 style.Set("bottom", "0")
}
