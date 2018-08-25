package app

type Style struct {
	css Css
}

func (style *Style) SetBackgroundColor(color string) {
	style.css.Set("background-color", color)
}

func (style *Style) SetLayout(layout string) {
	style.css.Set("display", layout)
}

func (style *Style) AutoExpand() {
	style.css.Set("flex-grow", "1")
}

func (style *Style) SetPosition(x, y string) {
	 style.css.Set("top", y)
	 style.css.Set("left", x)
}

func (style *Style) SetSize(width, height string) {
	 style.css.Set("width", width)
	 style.css.Set("height", height)
}

func (style *Style) SetMaxHeight(height string) {
	 style.css.Set("max-height", height)
}

func (style *Style) SetHeight(height string) {
	 style.css.Set("height", height)
}

func (style *Style) SetWidth(width string) {
	 style.css.Set("width", width)
}

func (style *Style) SetSticky() {
	 style.css.Set("position", "fixed")
}

func (style *Style) SetHidden() {
	 style.css.Set("display", "none")
}

func (style *Style) SetVisible() {
	 style.css.Set("display", "initial")
}

func (style *Style) AttachLeft() {
	 style.css.Set("left", "0")
}

func (style *Style) AttachBottom() {
	 style.css.Set("bottom", "0")
}
