package style

type Style struct {
	Css Css
}

func (style *Style) AnimatePosition(time, mode string) {
	style.Css.Set("transition", "top "+time+" ")
	style.Css.Set("transition", "left "+time+" ")
}


func (style *Style) SetBackgroundColor(color string) {
	style.Css.Set("background-color", color)
}

func (style *Style) SetLayout(layout string) {
	style.Css.Set("display", layout)
}

func (style *Style) SetPadding(padding string) {
	style.Css.Set("padding", padding)
}

func (style *Style) SetPaddingLeft(padding string) {
	style.Css.Set("padding-left", padding)
}

func (style *Style) SetFilter(filter string) {
	style.Css.Set("filter", filter)
}

func (style *Style) SetPaddingBottom(padding string) {
	style.Css.Set("padding-bottom", padding)
}

func (style *Style) CenterX() {
	style.Css.Set("display", "block")
	style.Css.Set("margin-right", "auto")
	style.Css.Set("margin-left", "auto")
	style.Css.Set("text-align", "center")
}


func (style *Style) AutoExpand() {
	style.Css.Set("flex-grow", "1")
}

func (style *Style) SetPosition(x, y string) {
	 style.Css.Set("top", y)
	 style.Css.Set("left", x)
}

func (style *Style) SetSize(width, height string) {
	 style.Css.Set("width", width)
	 style.Css.Set("height", height)
}

func (style *Style) SetMaxHeight(height string) {
	 style.Css.Set("max-height", height)
}

func (style *Style) SetDepth(depth string) {
	 style.Css.Set("z-index", depth)
}

func (style *Style) SetHeight(height string) {
	 style.Css.Set("height", height)
}

func (style *Style) SetWidth(width string) {
	 style.Css.Set("width", width)
}

func (style *Style) SetSticky() {
	 style.Css.Set("position", "fixed")
}

func (style *Style) SetHidden() {
	 style.Css.Set("display", "none")
}

func (style *Style) SetVisible() {
	 style.Css.Set("display", "block")
}

func (style *Style) Flex() {
	 style.Css.Set("display", "flex")
}

func (style *Style) Contain() {
	 style.Css.Set("object-fit", "contain")
}

func (style *Style) AttachTop() {
	 style.Css.Set("top", "0")
}

func (style *Style) AttachLeft() {
	 style.Css.Set("left", "0")
}

func (style *Style) AttachRight() {
	 style.Css.Set("right", "0")
}

func (style *Style) AttachBottom() {
	 style.Css.Set("bottom", "0")
}
